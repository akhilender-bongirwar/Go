package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files directory")
	content := "This is file 1 writing - we have to write in the first file"

	file, err := os.Create("./myFirst.txt")

	checkNilErr(err)

	len, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("The lenght is", len)

	defer file.Close()

	readFile("./myFirst.txt")
}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	// fmt.Println("The content is", databyte)
	fmt.Println("The content is", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
