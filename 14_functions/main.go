package main

import "fmt"

func main() {
	fmt.Println("Bro this is go dude where is your humor!!")
	first()
	res := mult(6, 7)
	fmt.Println(res)
	result, resultRemark := multi(6, 7, 8, 9, 10, 11, 13, 18)
	fmt.Println(result)
	fmt.Println(resultRemark)
}

func first() {
	fmt.Println("This is first funtion call")
}
func mult(val1 int, val2 int) int {
	return val1 * val2
}
func multi(values ...int) (int, string) {
	total := 1
	for _, val := range values {
		total *= val
	}
	return total, "Nice multiplication"
}
