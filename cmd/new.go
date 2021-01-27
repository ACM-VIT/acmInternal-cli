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
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func inputLine(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		auth.Check(err)
		return " "
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	return input
}

func newProject() {
	accessToken, err := auth.Login()

	fmt.Print("New Project Screen :\n\n")

	var name string

	var status = "ideation"

	//  fmt.Print("Project Name :");
	// _, errp := fmt.Scanln(&name);
	// auth.Check(errp);

	name = inputLine("Project Name :")

	var desc string

	// fmt.Print("Project Desc: ")
	// _, errq := fmt.Scanln(&name);
	// auth.Check(errq);

	desc = inputLine("Project Desc :")

	postBody, _ := json.Marshal(map[string]string{
		"name":   name,
		"desc":   desc,
		"status": status,
	})
	responseBody := bytes.NewBuffer(postBody)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", auth.BaseURL+"/v1/project/new", responseBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Bearer "+accessToken)
	resp, _ := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Check your internet connection")
		log.Fatalln(err)
	}
	// sb := string(body)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	//log.Printf(sb)
	if resp.Status != "200 OK" {
		fmt.Printf("\nerror:Unable to create Project: %v\n", resp.Status)
		auth.DumpMap("", data)
		os.Exit(1)
	}

	fmt.Println("Sucessfully created project !")

}

func newMeeting() {
	//accessToken,err := auth.Login();
	fmt.Println("new meeting")
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func newResource(projectName string) {
	accessToken, err := auth.Login()
	auth.Check(err)
	//	fmt.Println(accessToken);
	client := &http.Client{}
	req, _ := http.NewRequest("GET", auth.BaseURL+"/v1/project/fetch/byName/"+projectName, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Bearer "+accessToken)
	res, _ := client.Do(req)
	if res.Status != "200 OK" {
		fmt.Printf("\nerror:Unable to fetch Project: %v\n\ninfo: project may not exist or has been deleted from archives \n", projectName)
		os.Exit(1)
	}
	defer res.Body.Close()
	//buf is byte version of the json body
	buf, _ := ioutil.ReadAll(res.Body)

	//since the json is unstructered we use map
	var data map[string]interface{}
	err = json.Unmarshal([]byte(buf), &data)
	if err != nil {
		panic(err)
	}

	//dumpMap(" ", data)
	userInfo := data["data"].(map[string]interface{})
	project := userInfo["project"].([]interface{})
	for _, item := range project {
		id := item.(map[string]interface{})["id"].(string)
		projectName := item.(map[string]interface{})["name"].(string)
		fmt.Printf("New Resource Link for %v \n\n", projectName)
		resourceName := inputLine("Enter Resource Name: ")
		resourceLink := inputLine("Enter Resource Link: ")
		if !isURL(resourceLink) {
			fmt.Println("error: resource link must be a valid url")
			return
		}
		//fmt.Printf("%v \n %v \n %v \n", id, resourceName, resourceLink)
		//dumpMap(" ",item.(map[string]interface{}));
		//network request to update the project with the resource link
		//fmt.Printf("%v %v \n", resourceName, resourceLink)
		postBody2, _ := json.Marshal(map[string]string{
			resourceName: resourceLink,
		})
		reqBody2 := bytes.NewBuffer(postBody2)
		client2 := &http.Client{}
		req2, _ := http.NewRequest("PUT", auth.BaseURL+"/v1/project/update/projectResourcesLinks/"+id, reqBody2)
		//req2, _ := http.NewRequest("PUT", "https://webhook.site/90bd5377-0f66-43b7-8557-fa3bf41522a2", reqBody2)
		req2.Header.Set("Content-Type", "application/json")
		req2.Header.Set("authorization", "Bearer "+accessToken)
		res2, _ := client2.Do(req2)
		var data2 map[string]interface{}
		buf2, _ := ioutil.ReadAll(res2.Body)

		//since the json is unstructered we use map
		err2 := json.Unmarshal([]byte(buf2), &data2)
		if err2 != nil {
			panic(err2)
		}
		defer res2.Body.Close()
		//fmt.Println("status ", res2.Status)
		//dumpMap("", data2)
		if res2.Status != "200 OK" {
			fmt.Printf("\nerror:Unable to update project with links:")
			os.Exit(1)
		} else {
			fmt.Printf("\n success: updated project %v with resource links\n", projectName)
		}
	}

	//buf is byte version of the json body

}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "use to create a new project or new meeting or new link for a project",
	Long: ` 
	Takes 3 arguments types :
	acm new project [to generate a new project]
	acm new meeting [to generate a new meeting]
	acm new link (projectName) [to add a new resource link to a existing project]
`,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "project":
			newProject()
		case "meeting":
			newMeeting()
		case "link":
			if len(args) < 2 {
				fmt.Println("please enter a project name as your second argument: [acm new link (projectName)]")
			} else {
				newResource(args[1])
			}
		default:
			fmt.Println("invalid argument")
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
