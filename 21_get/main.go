package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("HELlo")
	link := "http://localhost:5050/postform"
	// HandleGetReq(link)
	HandlePostFormReq(link)
}

func HandleGetReq(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // Never forget this even use defer
	fmt.Println("The status code is ", res.StatusCode)
	fmt.Println("The Content lenght is ", res.ContentLength)
	//! One Way
	// content, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(content))

	//! The second way to do the same
	var responseString strings.Builder //? Just add this strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is:", byteCount)
	fmt.Println(responseString.String())
}
func HandlePostReq(url string) {
	reqBody := strings.NewReader(`
	 {
		"Subject":"I am learning go lang",
		"Description":"Want to get into devops"
	 }
	`)
	res, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
func HandlePostFormReq(myUrl string) {

	data := url.Values{}

	data.Add("firstName", "Akhil")
	data.Add("LastName", "B")
	data.Add("email", "akhilb@go.dev")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
