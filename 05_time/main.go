package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hey do you know what time it is?")
	fmt.Println("So the time is as follows...")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdTime := time.Date(2005, time.February, 24, 18, 30, 0, 0, time.UTC)
	fmt.Println(createdTime)
	fmt.Println(createdTime.Format("01-02-2006 15:04:05 Monday"))
}
