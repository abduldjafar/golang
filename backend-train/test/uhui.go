package main

import (
	"fmt"
	"time"
)
func main() {

	currentTime := time.Now()

	fmt.Println("Current Time in String: ", currentTime.String())

	fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
}