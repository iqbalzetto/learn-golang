package main

import (
	"errors"
	"fmt"
)

// It's possible to define custom error types by implementing
// the Error() method on them. Here's a variant on the exmaple
// above that uses a custom type to explicitly represent an argument error.

// A cusom error usually has the suffix "Error"
type argError struct {
	arg     int
	message string
}

// Adding this error method makes argError implement error interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// return custom error
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}

	}
	return arg + 3, nil
}

// errors.As is a more advanced version of errors.Is. It
// checks that a given error (or any error in it's chain)
// matches a specific error type and converts to a value of
// that type, returning true. If there's no match, it returns false.
func main() {
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
