/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
	"regexp"
	"os"
	"log"
	"golang.org/x/crypto/ssh/terminal"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/boltdb/bolt"
	"time"
	"reflect"
)

// type dataRespones struct {
// 	id string
// 	full_name string
// 	fcm_token string
// 	email string
// 	pwd string
// 	profilePic string
// 	verified bool
// 	accounts_connected:
// }

// type loginResponse struct{
//     statusCode int
//     message string
//     data {
//         "id": "FYgdYfHBlgMqclapvG8j",
//         "user": {
//             "id": "FYgdYfHBlgMqclapvG8j",
//             "full_name": "Gokul Kurup",
//             "fcm_token": "fdO1ANb-TzOtaeZ0Da8MGl:APA91bE4AR0TiZPXXin-8u8u9WZqtiUJgLDJ1g6S209Ptxse7vc987xcGKDGHNBmtVp7sNyMv_4LhIhyKZWfu6mrF2Evw-SoXn45RAQuBWDRfKD9JTMigZiKWHEZNSkVrvWw256sLXxi",
//             "email": "kurupgokul11@gmail.com",
//             "pwd": "$2a$10$NrxJDuYGyqRtglWDwdArx.cZ6CM8IkMWpZlqqZnYi7LTNaHXE8n5y",
//             "profilePic": "https://lh3.googleusercontent.com/a-/AOh14Gh_12XeBFf7nEXQYLIe6hh-OowXBURJzTi2js8Ohg=s96-c",
//             "verified": false,
//             "accounts_connected": {
//                 "discord": {
//                     "id": "373131099976499200",
//                     "username": "Madrigal1",
//                     "locale": "en-US",
//                     "flags": 0,
//                     "discriminator": "1465",
//                     "email": "kurupgokul11@gmail.com",
//                     "verified": true,
//                     "public_flags": 0,
//                     "mfa_enabled": true,
//                     "avatar": null
//                 },
//                 "google": {
//                     "pic": "https://lh3.googleusercontent.com/a-/AOh14Gh_12XeBFf7nEXQYLIe6hh-OowXBURJzTi2js8Ohg=s96-c",
//                     "name": "Gokul Kurup",
//                     "id": "106658813252341093333",
//                     "userid": "106658813252341093333",
//                     "email_verified": true,
//                     "email": "kurupgokul11@gmail.com"
//                 }
//             },
//             "discord_username": "Madrigal1"
//         },
//         "tokens": {
//             "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY212aXQuaW4iLCJhdWQiOiJhY212aXQuaW4iLCJzdWIiOiJGWWdkWWZIQmxnTXFjbGFwdkc4aiIsImlhdCI6MTYwNzM1NTkxNywiZXhwIjoxNjA5OTQ3OTE3LCJwcm0iOiIzMDA5M2NkNTExNTNlMTQ3YzkzZDY3N2FiOGQ3NjA1MmYyZTE2NDg2MWFhNmQwY2E2OTZjYWFiNzJmMzZiYTcyOTE0NDhkZWE2ZWU4MmY0MWZiMWE4NzhjYTNjY2QyOTBjY2EzOThmYzc2MTg3OWIzZTZhNWJiZjc5ZGFjOGVlZCJ9.BbNw8Wu-GtZ6QM2Tuz-P1S0fsdeSKFeVMMzvUvOtJyvH4ADJ0D6qDfeDPrfpFhCQTzSba84KN9YIU11Ua2D7VQ",
//             "refreshToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhY212aXQuaW4iLCJhdWQiOiJhY212aXQuaW4iLCJzdWIiOiJGWWdkWWZIQmxnTXFjbGFwdkc4aiIsImlhdCI6MTYwNzM1NTkxNywiZXhwIjoxNjE3NzIzOTE3LCJwcm0iOiI4ZGQ3NmJlNGE4MTMzMWI2YWMxYmMzYTBkMGQ5NjQzYjQ4Y2E4YzVkN2M1ODk1MDhhMTRlMzM4OWY4YWNmOWI1ZTFkOTdiNTc0M2M5NWZhOGFlNWRiNDI2MjI1YjcyMGZhYmRkZDAwMjUzN2VlNWJjODY2YmI5MjhhODZjNDU5ZiJ9.OzUiPsileiW0ddpKSGLZqfPIczf2Tapo6yZoFLJJ9Dv8TxIw2TgNLbvUsTo9ohCm4d-mbc9XjWGjmwpwUmDs1g"
//         }
//     }

//BaseURL for backend
var BaseURL string;

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$");
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func getPassword() string {
    // https://godoc.org/golang.org/x/crypto/ssh/terminal#ReadPassword
    // terminal.ReadPassword accepts file descriptor as argument, returns byte slice and error.
    passwd, e := terminal.ReadPassword(int(os.Stdin.Fd()))
    if e != nil {
        log.Fatal(e)
    }
    // Type cast byte slice to string.
    return string(passwd)
}

func saveTokenToDb(accessToken string)  {
	db, err := bolt.Open("my.db",0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("Tokens"))
		err = b.Put([]byte("accessToken"), []byte(accessToken))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tokens"))
		v := b.Get([]byte("accessToken"))
		fmt.Printf("The at from db is: %s\n", v)
		return nil
	})
	defer db.Close()
}


func longRunningTask(accessToken string) <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)
		
		// Simulate a workload.
		saveTokenToDb(accessToken)
	}()

	return r
}


func loginReq(email string,pwd string) {
	postBody, _ := json.Marshal(map[string]string{
		"email": email,
		"pwd":pwd,
	 })
	 responseBody := bytes.NewBuffer(postBody)
  //Leverage Go's HTTP Post function to make request
	if BaseURL == " " || len(BaseURL)<2 {
		log.Fatalf("Base Url not found %v", BaseURL)
		os.Exit(1);
	}
	 resp, err := http.Post(BaseURL + "/v1/access/login/basic", "application/json", responseBody)
  //Handle Error
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
		fmt.Printf("error:Incorrect email or password \n")
		os.Exit(1)
	 } 
	//else process to put accessToken to db
	 userInfo := data["data"].(map[string]interface{})
		tokens := userInfo["tokens"].(map[string]interface{})
		accessToken:= tokens["accessToken"].(string)
		refreshToken:= tokens["refreshToken"].(string)
		 fmt.Printf("accessToken: %s \n refreshToken: %s" ,accessToken,refreshToken)
		 r := <-longRunningTask(accessToken)
		 if reflect.TypeOf(r).Kind() == reflect.Int {
			fmt.Printf("msg: Login Successful !\n")
		 }
		
	
}
func userLogin(args []string) {
	fmt.Printf("Enter Your Email: ") 
  
    // var then variable name then variable type 
    var email string 
  
    // Taking input from user 
	fmt.Scanln(&email) 
	if !isEmailValid(email) {
		fmt.Println("not a valid email")
		os.Exit(1);
	}
	var pwd string 
	fmt.Printf("Enter Password: ")
    pwd = getPassword()
    
    loginReq(email,pwd)
 }

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		userLogin(args);
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
