package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hey")
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.April, 12, 23, 23, 0, 0)
	fmt.Println(createdDate)
}
