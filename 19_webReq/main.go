package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://akhils-profile.vercel.app/"

func main() {
	fmt.Println("Hello now I will handle web requests")
	// This is how we write a get request
	res, err := http.Get(url)
	handleError(err)
	fmt.Printf("The response type is %T\n", res)
	defer res.Body.Close() // This is our responsibility to close the connection

	dataByte, err := ioutil.ReadAll(res.Body)
	handleError(err)
	content := string(dataByte)
	fmt.Println(content)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
