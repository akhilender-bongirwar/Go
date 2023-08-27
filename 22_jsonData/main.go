package main

import (
	"encoding/json"
	"fmt"
)

// ? Remember you have to capitalize the first letter of struct to export the fieldNames
// For suppose I put interest instead of Interests then it was not in ouptut
type studentDetails struct {
	Name     string `json:"studentname"`
	Branch   string
	RollNum  int      `json:"rollnum"`
	GPA      int      `json:"-"`
	Interest []string `json:"interests,omitempty"`
}

// ! This is where I can use `json:"{name}"` to get into json output and json"-" to remove it from the response
// ! and omitempty to remove spaces

func main() {
	fmt.Println("This is where we are converting all to json")
	encodeJson()
}

func encodeJson() {
	iiitl := []studentDetails{
		{"Akhil", "CS", 11, 9, []string{"web-dev", "software_development"}},
		{"Ram", "IT", 11, 10, []string{"app-dev", "software_development"}},
		{"Krsna", "CSAI", 11, 10, []string{"ml", "software_development"}},
		{"Shyam", "CSB", 11, 9, nil},
	}

	finalJson, err := json.MarshalIndent(iiitl, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
