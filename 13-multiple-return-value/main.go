/**
Go has built-in support for multiple return values. This feature is used often in
idiomatic Go, for example to return both result and error values from a function.
*/

package main

import "fmt"

// the (int, int)  in this function signature shows that the function returns 2 ints
func vals() (int, int) {
	return 3, 7
}

//here we use the 2 different return values from the call with multiple assignment

//if you only want a subset of the returned values, use the blank identifier _

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, c := vals()
	fmt.Println(c)
}
