package main

import "fmt"

func main() {
	defer fmt.Println("World!!")
	defer fmt.Println("First!!")
	defer fmt.Println("second!!")
	fmt.Println("Hello - lastLine")
	myDefer()
}

// basically stack last in first out

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
