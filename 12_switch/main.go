package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1 and you move 1")
		fallthrough
	case 2:
		fmt.Println("Dice value is 2 and now you move 2")
		fallthrough
	case 3:
		fmt.Println("Dice val is 3 and so on ....")
		fallthrough
	default:
		fmt.Println("Handle it yourselves bro!!")
	}
}
