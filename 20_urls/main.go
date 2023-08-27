package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=jerjwjfjirhe"

func main() {
	fmt.Println("This is handling urls")
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)
	// // fmt.Println(result)
	// fmt.Printf("The type of result is %T", result)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Path)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of result is %T \n", qparams)
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param value is: ", val)
	}

	partOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	newUrl := partOfUrl.String()
	fmt.Println(newUrl)

}
