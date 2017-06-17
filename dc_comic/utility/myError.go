package utility

import "fmt"

// MyError is an error structure
type MyError struct {
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v", e.What)
}
