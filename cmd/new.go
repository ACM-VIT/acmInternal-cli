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
	"errors"
	"github.com/spf13/cobra"
	auth "acm/cli-core"
	"encoding/json"
	"bytes"
	"net/http"
	"log"
	"io/ioutil"
	"os"
)


func isValidCommand(command string) bool {
	switch command {
	case
		"project",
		"meeting":
		return true
	}
	return false
}


func newProject() {
	accessToken,err := auth.Login();
	
	fmt.Printf("New Project :\n\n");

	var name string;
	var desc string;
	var status = "ideation";


	postBody, _ := json.Marshal(map[string]string{
		"name": name,
		"desc":desc,
		"status":status,
	 })
	 responseBody := bytes.NewBuffer(postBody);
	 client := &http.Client{};
	req, _ := http.NewRequest("POST",auth.BaseURL+"/v1/project/fetch/all",responseBody);
	req.Header.Set("authorization","Bearer " + accessToken);
	resp, _ := client.Do(req)
	 if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()
  //Read the response body
	 body, err := ioutil.ReadAll(resp.Body)
	 if err != nil {
		log.Printf("Check your internet connection");
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
		fmt.Printf("\nerror:Incorrect email or password \n")
		os.Exit(1)
	 } 

	 fmt.Println("Sucessfully created project !")

}

func newMeeting() {
	//accessToken,err := auth.Login();
	fmt.Println("new meeting");
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "use to create a new project or new meeting",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
		return errors.New("requires a argument. [either project or meeting]")
		}
		
		if(!isValidCommand(args[0])) {
			return errors.New("invalid argument. [must be project or meeting]");
		}
		return nil;
  	},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
			case "project": newProject();
			case "meeting": newMeeting();
			default: fmt.Println("invalid argument");	
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
