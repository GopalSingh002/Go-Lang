/*
We can remove the condition from the switch case, and use different condition for each case
it help us to replace the if-then-else chain
ex:
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("Good morning")
	case t < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good Evenning")
	}

}
