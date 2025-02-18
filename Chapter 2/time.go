/*
time package allow us to use date, time features in our program.
time in go playground always start at 2009-11-10 23:00:00 UTC
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())           // current date, time
	fmt.Println(time.Now().Weekday()) // current week day
	fmt.Println(time.Now().Date())    //current date
	fmt.Println(time.Now().Day())     // current day-date
	fmt.Println(time.Now().Month())   //current month
	fmt.Println(time.Now().Year())    //current year
	fmt.Println(time.Now().Hour())    // current hour
	fmt.Println(time.Now().Minute())  //current min

}
