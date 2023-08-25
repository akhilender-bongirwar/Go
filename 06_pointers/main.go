package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNum := 89
	ptr := &myNum
	fmt.Println("The address of myNum is", ptr)
	fmt.Println("The address of myNum is", *ptr)

}
