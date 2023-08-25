package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("I am slice. You love to eat me")
	var fruitList = []string{"Hi"}
	fmt.Printf("the fruitlist is of %T ", fruitList)
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	fruitList = append(fruitList, "Bro do ", "you think this is ", "a fruitList")
	fmt.Println(len(fruitList))
	fmt.Println(fruitList)

	fruitList = append(fruitList[:1])
	fmt.Println(len(fruitList))
	fmt.Println(fruitList)

	marks := make([]int, 2)
	marks[0] = 100
	marks[1] = 78
	fmt.Println(marks)
	marks = append(marks, 67, 39, 10)
	fmt.Println(marks)
	sort.Ints(marks)
	fmt.Println(sort.IntsAreSorted(marks))
	fmt.Println(marks)

	// removing a certain index in go
	index := 2
	var courses = []string{"react", "js", "swift", "c++", "ruby", "go"}
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
