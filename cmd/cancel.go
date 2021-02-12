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
	"os"
	"github.com/spf13/cobra"
	auth "acm/cli-core"
	"net/http"
	"bytes"
	"encoding/json"
)

func cancelMeeting() {
	accessToken, err := auth.Login()
	auth.Check(err)
	//	fmt.Println(accessToken);

	var title string;
	title = inputLine(`Input title of meeting: `);

	type Request struct {
		Title string `json:"title"`
	}
	postBody, _ := json.Marshal(Request{
		Title:   title,
	})
	responseBody := bytes.NewBuffer(postBody)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", auth.BaseURL+"/v1/meeting/cancel", responseBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Bearer "+accessToken)
	res, _ := client.Do(req)
	if res.Status != "200 OK" {
		fmt.Printf("error:failed to delete/cancel meeting")
		os.Exit(1)
	} else {
		fmt.Println("sucess: Successfully cancelled Meeting in db and google calender")
	}
	defer res.Body.Close();
	return;
}

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "cancel your meeting using title",
	Long: "usage:\nType: acm cancel (meetingTitle)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("error: invalid argument:type acm help cancel")
			os.Exit(1)
		} else {
			cancelMeeting();
		}
	},
}

func init() {
	rootCmd.AddCommand(cancelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
