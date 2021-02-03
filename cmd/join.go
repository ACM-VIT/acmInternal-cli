/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	auth "acm/cli-core"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func joinProject(projectName string) {
	accessToken, err := auth.Login()
	auth.Check(err)
	//	fmt.Println(accessToken);
	client := &http.Client{}
	req, _ := http.NewRequest("GET", auth.BaseURL+"/v1/project/fetch/byName/"+projectName, nil)
	req.Header.Set("authorization", "Bearer "+accessToken)
	res, _ := client.Do(req)
	if res.Status != "200 OK" {
		fmt.Printf("\nerror:Unable to fetch Project: %v\n\ninfo: project may not exist or has been deleted from archives \n", projectName)
		os.Exit(1)
	}
	buf, _ := ioutil.ReadAll(res.Body)

	//since the json is unstructered we use map
	var data map[string]interface{}
	err = json.Unmarshal([]byte(buf), &data)
	if err != nil {
		panic(err)
	}
	userInfo := data["data"].(map[string]interface{})
	project := userInfo["project"].([]interface{})
	for _, item := range project {
		projectID := item.(map[string]interface{})["id"].(string)
		client2 := &http.Client{}
		req2, _ := http.NewRequest("PUT", auth.BaseURL+"/v1/project/join/"+projectID, nil)
		req2.Header.Set("authorization", "Bearer "+accessToken)
		res2, _ := client2.Do(req2)
		if res2.Status != "200 OK" {
			fmt.Printf("\nerror:Unable to add you as Team Member: %v\n\ninfo: mostly internal error with backend :(. \nPlease try again with good internet or rant/abuse maintainers \n", projectName)
			os.Exit(1)
		} else {
			fmt.Printf("success: Successfully joined project: " + projectName + "\n")
		}
		//dumpMap(" ",item.(map[string]interface{}));
	}
}

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Use command to join a project",
	Long: `
	Usage:
		acm join (projectname)
		
	tip: after a sucess message use acm projects (projectName)
		  to see detailed info`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(" error : command requires the name of the project you wanna join as argument.\n put in quotes if it contains a space or more than one word")
		} else {
			joinProject(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(joinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
