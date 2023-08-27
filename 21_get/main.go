package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("HELlo")
	link := "http://localhost:5050/post"
	// HandleGetReq(link)
	HandlePostReq(link)
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
