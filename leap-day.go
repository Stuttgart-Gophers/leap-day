package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.UnixDate))

	currentMonth := now.Local().Month()
	currentDay := now.Day()

	if currentMonth.String() == "February" && currentDay == 29 {
		fmt.Println("Happy Leap Day!")
	} else {
		fmt.Println("Today is not a leap day.")
	}

}
