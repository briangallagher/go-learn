package string

import (
	"fmt"
)

// Uppercase means it is exported and can be used by other packages
func Reverse(s string) string {
	fmt.Println("Reversing....", s)
	return "Reversing...."
}
