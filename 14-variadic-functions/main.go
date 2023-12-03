/**
Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a
common variadic function.
*/

package main

import "fmt"

//Here's a function that will take an arbitrary number of ints as arguments.
//within the function, the type of nums is equivalent to []int. We can call
//len(nums) over it with range etc.

func sum(nums ...int) {
	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

//variadic functions can be called in the usual way with individual arguments

//if you already have multiple args in a slice, apply them to a variadic function
//using func(slice...) like this.

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
