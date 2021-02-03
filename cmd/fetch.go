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

func displayProjectsByTag(tagName string) {
	accessToken, err := auth.Login()
	auth.Check(err)
	//	fmt.Println(accessToken);
	client := &http.Client{}
	req, _ := http.NewRequest("GET", auth.BaseURL+"/v1/project/fetch/byTag/"+tagName, nil)
	req.Header.Set("authorization", "Bearer "+accessToken)
	res, _ := client.Do(req)
	fmt.Println("Projects by Acm")
	if res.Status != "200 OK" {
		fmt.Printf("\nerror:Unable to find any Projects of this tag \n")
		os.Exit(1)
	}
	//buf is byte version of the json body
	buf, _ := ioutil.ReadAll(res.Body)

	//since the json is unstructered we use map
	var data map[string]interface{}
	err = json.Unmarshal([]byte(buf), &data)
	if err != nil {
		panic(err)
	}
	userInfo := data["data"].(map[string]interface{})
	allProjects := userInfo["allProjects"].([]interface{})
	for index, item := range allProjects {
		name := item.(map[string]interface{})["name"]
		desc := item.(map[string]interface{})["desc"]
		status := item.(map[string]interface{})["status"]
		founder := item.(map[string]interface{})["founder"]
		founderName := founder.(map[string]interface{})["full_name"]

		fmt.Printf("%d.\n Name: %v \n Desc: %v \n Status: %v \n Maintainer: %v  \n\n", index+1, name, desc, status, founderName)
		//dumpMap(" ",item.(map[string]interface{}));
	}

}

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "used to get project by Filter [tag]",
	Long: `Usage:
		acm fetch tag (tagName)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("invalid argument: \nType: acm help fetch ")
			return
		}
		switch args[0] {
		case "tag":
			if len(args) < 2 {
				fmt.Println("please enter a project name as your second argument: [acm fetch tag (projectName)]")
			} else {
				displayProjectsByTag(args[1])
			}
		default:
			fmt.Println("invalid argument: \nType: acm help new ")
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
