package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// GetToken : Struct for token
type GetToken struct {
	Token string `json:"token"`
}

func main() {
	// initiate the client
	client := &http.Client{}
	// set the form
	form := url.Values{}
	form.Add("userKey", "xxxxxxxxxxxxxxxxxxxxxxxxxxx")
	form.Add("userSecret", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	// create a  new request
	req2, err2 := http.NewRequest("POST", "https://api.s3zipper.com/gentoken", strings.NewReader(form.Encode()))
	if err2 != nil {
		log.Fatal("NewRequest: ", err2)
	}
	// set form in request
	req2.PostForm = form
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// get response
	resp3, err3 := client.Do(req2)
	if err3 != nil {
		log.Fatal("NewRequest: ", err3)
	}
	defer resp3.Body.Close()
	// read the body
	body, err4 := ioutil.ReadAll(resp3.Body)
	if err4 != nil {
	}

	/*******************************************************************/
	// Get the token from the body
	var p GetToken
	err := json.Unmarshal(body, &p)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}
	token := p.Token
	fmt.Println(token)

}
