package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	/**************** GET TOKEN ***************************/
	token := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	var bearer = "Bearer " + token

	/**************** GET UUIDS FROM ZIPSTART ***************************/
	allBodyUUIDs := []byte(`{
		"message": "STARTED",
		"size": "9.1 kB",
		"chainTaskUUID": [
		 {
		  "idurl": "task_7ad8b9e3-6d63-40a2-b7d7-7d337daf"
		 },
		 {
		  "email": "task_354596bd-66a9-42f4-83b0-c86eeb01"
		 }
		]
	   }`)
	/**************** ACCESS API WITH TOKEN  ***************************/
	client := &http.Client{}
	// create new request with allbody set
	//allbody contains uuids
	req2, err2 := http.NewRequest("POST", "https://api.s3zipper.com/v1/zipresult", strings.NewReader(string(allBodyUUIDs)))
	if err2 != nil {
		log.Fatal("NewRequest: ", err2)
	}
	//req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Set("Content-Type", "application/json; charset=utf-8")
	req2.Header.Set("Authorization", bearer)
	// get response
	resp3, err3 := client.Do(req2)
	if err3 != nil {
		log.Fatal("NewRequest: ", err3)
	}
	defer resp3.Body.Close()
	// read body
	body2, err4 := ioutil.ReadAll(resp3.Body)
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(string(body2[:]))

}
