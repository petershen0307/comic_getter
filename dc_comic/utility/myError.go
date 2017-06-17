package utility

import "fmt"

// MyError is an error structure
type MyError struct {
	what string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v", e.what)
}
