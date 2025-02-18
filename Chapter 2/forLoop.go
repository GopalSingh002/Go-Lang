/*
Go allow only for loop
Reasons:
- Simple, predictable
- Easily mimic the functioning of while, do-while

About:
for loop in go consists :
- init: executed before first iteration(start)
- condition: evaluate before each iteration
- post : executed after each iteration

Syntax:
for init; condition; post{
	code
}

Note:
- variables in the for visible only in the scope of loop.
- loop stop when the condition became false

	ex:
*/

package main

import "fmt"

func main() {
	//square till 10
	for i := 2; i <= 10; i++ {
		fmt.Println(i * i)
	}

	//init and post statment are optional,  even we dont need to use the the semicolons for post and init
	sq := 2
	for sq <= 10 {
		fmt.Println(sq * sq)
		sq += 1 // we need to mention this, so that loop can reach failing point
	}

	//if we do not give anything in the loop, it will run for infinite times.
	/*
		for{
		}
	*/

}
