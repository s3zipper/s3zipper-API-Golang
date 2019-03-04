package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// GetToken : Struct for token
type GetToken struct {
	Token string `json:"token"`
}

func main() {
	resp, err := http.PostForm("https://api.s3zipper.com/gentoken",
		url.Values{"userKey": {"xxxxxxxxxxxxxxxxxxxxxx"}, "userSecret": {"xxxxxxxxxxxxxxxxxxxxxxxxx"}})

	if err != nil {
		fmt.Println("errorination happened getting the response", err)
	}
	defer resp.Body.Close()
	/*******************************************************************/
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("\nbody -- %v\n", string(body))

	if err != nil {
		fmt.Println("Error happened reading the body", err)
	}
	/*******************************************************************/
	var p GetToken
	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Println("NewRequest: ", err)
	}
	token := p.Token
	//fmt.Println(map[string]string{"token": token})
	fmt.Println(token)

}
