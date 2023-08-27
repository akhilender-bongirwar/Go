package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("HELlo")
	link := "http://localhost:5050"
	HandleGetReq(link)
}

func HandleGetReq(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // Never forget this
	fmt.Println("The status code is ", res.StatusCode)
	fmt.Println("The Content lenght is ", res.ContentLength)
	//! One Way
	// fmt.Println(string(content))

	//! The second way to do the same
	var responseString strings.Builder //? Just add this strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is:", byteCount)
	fmt.Println(responseString.String())
}
