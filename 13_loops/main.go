package main

import "fmt"

func main() {
	fmt.Println("Hello World!!")
	// days := []string{"Mon", "Tue", "Thur", "Fri", "Sat", "Sun"} // this is a slice I have created
	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	//! Another format
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }
	//? Another Format
	// for _, val := range days {
	// 	fmt.Printf("The elements of the slices are %v \n", val)
	// }
	value := 3
	for value < 10 {
		if value == 5 {
			goto label
		}
		fmt.Println("Value is:", value)
		value++
	}

label:
	fmt.Println("Bro you are my friend!!")
}
