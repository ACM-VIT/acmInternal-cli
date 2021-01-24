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
	"fmt"
	auth "acm/cli-auth"
	"github.com/spf13/cobra"
	"net/http"
	"io/ioutil"
	//"github.com/buger/jsonparser"
	"encoding/json"
	"os"
)

type project struct{
	id string
	name string
	desc string
	status string
}


func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}


func displayProjects() {
	accessToken,err := auth.Login();
	auth.Check(err);
//	fmt.Println(accessToken);
	client := &http.Client{};
	req, _ := http.NewRequest("GET",auth.BaseURL+"/v1/project/fetch/all",nil);
	req.Header.Set("authorization","Bearer " + accessToken);
	res, _ := client.Do(req)
	fmt.Println("Projects by Acm");
	if res.Status != "200 OK" {
		fmt.Printf("\nerror:Unable to fetch Projects\n")
		os.Exit(1)
	 } 
	//buf is byte version of the json body
	buf, _ := ioutil.ReadAll(res.Body);

	//since the json is unstructered we use map
	var data map[string]interface{}
	err = json.Unmarshal([]byte(buf), &data)
	if err != nil {
		panic(err)
	}
	userInfo := data["data"].(map[string]interface{});
	allProjects := userInfo["allProjects"].([]interface{})
	for index,item :=range allProjects {
		name:= item.(map[string]interface{})["name"];
		desc:= item.(map[string]interface{})["desc"];
		status:= item.(map[string]interface{})["status"];
		fmt.Printf("%d.\n Name: %v \n Desc: %v \n Status: %v \n\n",index+1,name,desc,status);
		//dumpMap(" ",item.(map[string]interface{}));
	}

	
}

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "get all projects or search for a single project by adding it as a argument",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		displayProjects();
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
