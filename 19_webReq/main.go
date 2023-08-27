package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://akhils-profile.vercel.app/"

func main() {
	fmt.Println("Hello now I will handle web requests")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("The response type is %T\n", res)
	defer res.Body.Close() // This is our responsibility to close the connection

	dataByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataByte)
	fmt.Println(content)
}
