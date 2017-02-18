package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	const timeFormat = "01/02/2006"
	arg := os.Args[1]
	t, err := time.Parse(timeFormat, arg)
	d := t.Weekday().String()
	if err != nil {
		fmt.Println("The entered date is not valid.")
	} else {
		t.Format(timeFormat)
		fmt.Printf("The Weekday for %v is %v\n", t.Format(timeFormat), d)
	}
}
