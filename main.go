package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	const timeFormat = "01/02/2006"
	arg := os.Args[1]
	t, _ := time.Parse(timeFormat, arg)
	day := t.Weekday().String()

	fmt.Printf("The Weekday for %v %v %v is %v\n", t.Month(), t.Day(), t.Year(), day)
}
