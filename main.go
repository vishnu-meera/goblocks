package main

import (
	"fmt"
)

func main() {
	input := "Hello\n\nWorld\nThis is Go"

	for idx, line := range lines {
		fmt.Printf("Line %d: %q\n", idx+1, line)
	}
}
