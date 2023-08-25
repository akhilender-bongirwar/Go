package main

import "fmt"

func main() {
	fmt.Println("These are the structure in go")
	akhil := User{"Akhil", "I am cool", true, 1}
	fmt.Println(akhil)
	fmt.Printf("akhil details are %+v \n", akhil)
	akhil.GetStatus()
	akhil.NewMsg()
	fmt.Println("Message is ", akhil.Message)
}

type User struct {
	Name    string
	Message string
	Status  bool
	Count   int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}
func (u User) NewMsg() {
	u.Message = "Bro wtf"
	fmt.Println(u.Message)
}
