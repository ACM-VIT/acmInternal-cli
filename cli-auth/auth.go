package auth

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"regexp"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
	"github.com/inancgumus/screen"
)

const (
	//BaseURL is the Base Url for the backend
	BaseURL = "https://us-central1-acminternal.cloudfunctions.net/App"
)

func Check(e error) {
    if e != nil {
        panic(e)
    }
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

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$");
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}


func Login()(string, error) {
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
	acessToken,err := getUserInfo(email,pwd);
	Check(err);
	return acessToken,err;
}

func getUserInfo(email string,pwd string)(string,error) {
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
		fmt.Printf("\nerror:Incorrect email or password \n")
		os.Exit(1)
	 } 
	//else process to put accessToken to db
	 userInfo := data["data"].(map[string]interface{})
		tokens := userInfo["tokens"].(map[string]interface{})
		accessToken:= tokens["accessToken"].(string)
		//refreshToken:= tokens["refreshToken"].(string)
		// fmt.Printf("accessToken: %s \n refreshToken: %s" ,accessToken,refreshToken)
		screen.Clear()
		screen.MoveTopLeft()
	return accessToken,nil;
}