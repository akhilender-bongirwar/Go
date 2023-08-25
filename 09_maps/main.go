package main

import "fmt"

func main() {
	fmt.Println("These are the maps")
	language := make(map[string]string)
	language["JS"] = "JavaScript"
	language["C++"] = "c++"
	language["RB"] = "Ruby"
	language["Py"] = "Python"
	delete(language, "C++")
	fmt.Println(language)

	for _, val := range language {
		fmt.Printf("For key , the value is %v \n", val)
	}

}
