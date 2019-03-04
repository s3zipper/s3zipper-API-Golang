package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	/**************** GET TOKEN FROM COOKIE OR SESSION ***************************/
	token := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	filePaths := []string{"{{bucket}}/files/pigeon.jpg", "{{bucket}}/files/file2.txt"}

	/**************** ACCESS API WITH TOKEN  ***************************/
	var bearer = "Bearer " + token
	client := &http.Client{}
	////////////////////////////////////////////////////////////////////
	form := url.Values{}
	form.Add("awsKey", "xxxxxxxxxxxxxxxxxxxxxxxx")
	form.Add("awsSecret", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	form.Add("awsBucket", "mybucket")
	form.Add("awsRegion", "us-east-1")
	//form.Add("awsToken", "")
	//form.Add("awsEndpoint", "")
	//form.Add("zipTo", "")
	form.Add("resultsEmail", "email@gmail.com") // email to send results to
	for _, filepath := range filePaths {        // Add filepaths here
		form.Add("filePaths", filepath)
	}
	//////////////////////////////////////////////////////////////////////////////////////////////////
	req2, err2 := http.NewRequest("POST", "https://api.s3zipper.com/v1/streamzip", strings.NewReader(form.Encode()))
	if err2 != nil {
		log.Fatal("NewRequest: ", err2)
	}

	req2.PostForm = form
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Set("Authorization", bearer)
	///////////////////////
	resp3, err3 := client.Do(req2)
	if err3 != nil {
		log.Fatal("NewRequest: ", err3)
	}
	defer resp3.Body.Close()
	//
	body2, err4 := ioutil.ReadAll(resp3.Body)
	if err4 != nil {
		fmt.Println(err4)
	}

	/****************** Return raw json *************************************************/
	fmt.Println(string(body2[:]))

}
