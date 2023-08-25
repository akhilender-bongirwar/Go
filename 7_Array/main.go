package main

import "fmt"

func main() {
	fmt.Println("These are the Arrays")

	var fruitList [4]string
	fruitList[0] = "I"
	fruitList[2] = "j"
	fruitList[3] = "k"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [4]string{"I", "J"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))

	var noVeg = []string{"I", "", "J"}
	fmt.Println(noVeg)
	fmt.Println(len(noVeg))

}
