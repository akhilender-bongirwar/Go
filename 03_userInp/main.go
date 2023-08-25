package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Hello welcome to user Input Learning"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name to continue:")
	// comma ok, syntax
	inp, _ := reader.ReadString('\n')
	fmt.Println("Hey, Nice meeting you ", inp)
}
