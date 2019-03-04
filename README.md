# S3zipper-API
S3zipper is a managed zipping service for Amazon S3.  
It is a lightweight API but robust in its capabilities.
It can handle zipping many Gigabytes of data efficiently.

# Documentation
1. Zipping to Amazon S3 bucket:  
[Zip to S3 API](https://docs.s3zipper.com/#23fc2566-464e-bcf7-1e0d-614dd77290df)
2. Stream zipping while downloading:  
[Stream S3 Downloads API](https://docs.s3zipper.com/#1c290c02-8c67-14d7-6fee-3912dca4abbf)
3. EC2 AMI only (Only for AWS EC2)  
[EC2(AMI) S3ZIPPER API SERVER](https://docs.s3zipper.com/#bd260c71-5f11-4a05-a07b-6e489ca8cb7d)

# Website
- Main Website:  
[S3zipper](https://s3zipper.com/)

- AWS EC2 AMI :  
[Amazon EC2](https://aws.amazon.com/marketplace/pp/B0727QDVXV)

# USAGE

## 1. Register for Account
``` URL : https://s3zipper.com/registration/login ```

Register for a new account or login to start the process.  
[Registration](https://s3zipper.com/registration/login)

## 2. Get credentials
``` URL : https://s3zipper.com/auth/develop ```  

We will need these credentials to later get tokens.  
[Developer](https://s3zipper.com/auth/develop)

## 3. Generate token
```API : https://api.s3zipper.com/gentoken```  

We will need tokens to securely access the rest of the API. Please save this token in a cookie or a session depending on your use case.

- Tokens are generated using credentials from step 2 above.

- All tokens last for 24 hours.

```go
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


```

## 4. Start zipping
The API currently supports two generally desired modes:  
### 4a). Zip to Amazon S3 bucket  
 ``` API : https://api.s3zipper.com/v1/zipstart  ```    

**zipstart:**  will zip files back into the same originating bucket and issue a download URL when done.

```go
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

	token := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	filePaths := []string{"{{bucket}}/files/file.jpg", "{{bucket}}/files/file2.txt"}
	/**************** ACCESS API WITH TOKEN  ***************************/
	var bearer = "Bearer " + token
	client := &http.Client{}
	////////////////////////////////////////////////////////////////////
	form := url.Values{}
	form.Add("awsKey", "xxxxxxxxxxxxxxxxxxxxxxxxxxx")
	form.Add("awsSecret", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	form.Add("awsBucket", "bucket")
	form.Add("awsRegion", "us-east-1")
	//form.Add("awsToken", "")
	//form.Add("awsEndpoint", "")
	//form.Add("zipTo", "")
	form.Add("resultsEmail", "email@gmail.com") // email to send results to
	for _, filepath := range filePaths {
		form.Add("filePaths", filepath)
	}
	//////////////////////////////////////////////////////////////////////////////////////////////////
	req2, err2 := http.NewRequest("POST", "https://api.s3zipper.com/v1/zipstart", strings.NewReader(form.Encode()))
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

	/****************** Return json *************************************************/
	fmt.Println(string(body2[:]))

}


```

### 4b). Stream download zip files on a browser
``` API :  https://api.s3zipper.com/v1/streamzip ```

 **streamzip:**  will generate a URL that can later be used to stream download files on a browser.

 ```go
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


 ```

## 5. Check your progess.
``` API : https://api.s3zipper.com/v1/zipstate ```

Some jobs can take quite a bit, and you might want to know their progress.  With this API call, you will get to know if the process completed successfully or if it is still running.
- For this, we will need the result from step 4.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	/**************** GET TOKEN FROM COOKIE ***************************/
	token := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	var bearer = "Bearer " + token

	/**************** GET UUIDS FROM Zipstart ***************************/

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
	// allbody contains uuids

	req2, err2 := http.NewRequest("POST", "https://api.s3zipper.com/v1/zipstate", strings.NewReader(string(allBodyUUIDs)))
	if err2 != nil {
		log.Fatal("err2: ", err2)
	}
	req2.Header.Set("Authorization", bearer)
	req2.Header.Set("Content-Type", "application/json; charset=utf-8")

	// get response
	resp3, err3 := client.Do(req2)
	if err3 != nil {
		log.Fatal("err3: ", err3)
	}
	defer resp3.Body.Close()
	// read body
	body2, err4 := ioutil.ReadAll(resp3.Body)
	if err4 != nil {
		log.Fatal("err4: ", err3)
	}
	fmt.Println(string(body2[:]))

}


```

## 6. Get the results.  
``` API : https://api.s3zipper.com/v1/zipresult```  

The API provides a background task that just listens and waits for the result. When done, it returns the result.   
**NB:**  
- It only returns the last result. If you zipped and requested the results emailed to you ; you will get a result about the email being sent.
- This is good if you need to automate things a bit. A good example is if you need to wait and send customized emails containing the result.
- This also consumes the result from step 4.

```go
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


```

## Conlusion

There it is! This API is more than three years in development, and we are putting a lot of effort into it to make sure that it gets even better. Our hope is to make it a de facto zipping service for busy people.

For now, it is intended to be a simple and relatively cheap API that just works and makes things easy.

More examples with different programming languages are available in [Documentation](https://docs.s3zipper.com/)
