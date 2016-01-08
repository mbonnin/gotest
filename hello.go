package main

import (
	"fmt"
)

// This is a comment
func main() {
	fmt.Printf("hello, world\n")

	conversion()
}

func conversion() {
	i := 42
	f := float64(42)
	u := uint(f)

	fmt.Println(i, f, u)

}
