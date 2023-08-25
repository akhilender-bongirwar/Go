package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Hello please enter your lucky number to start"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')

	num, err := strconv.ParseFloat(strings.TrimSpace(inp), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("I have increase 1 in your Lnumber ", num+1)
	}

}
