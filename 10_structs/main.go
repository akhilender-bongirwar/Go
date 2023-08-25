package main

import "fmt"

func main() {
	fmt.Println("These are the structure in go")
	akhil := User{"Akhil", "I am cool", true, 1}
	fmt.Println(akhil)
	fmt.Printf("akhil details are %+v", akhil)
}

type User struct {
	Name    string
	message string
	status  bool
	count   int
}
