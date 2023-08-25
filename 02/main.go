package main

import "fmt"

const LoginToken string = "65tre78"

// Public variable starts with capital Letter

func main() {
	var username string = "Akhil"
	fmt.Println(username)
	fmt.Printf("The type of variable is %T \n", username)

	var isbool bool = true
	fmt.Println(isbool)
	fmt.Printf("The type of variable is %T \n", isbool)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("The type of variable is %T \n", smallval)

	var largVal float32 = 255.5533434
	fmt.Println(largVal)
	fmt.Printf("The type of variable is %T \n", largVal)

	fmt.Println(LoginToken)
}
