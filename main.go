package main

import (
	"fmt"

	"github.com/vishnu-meera/goblocks/Utils"
)

func main() {
	input := "Hello\n\nWorld\nThis is Go"
	lines := Utils.SplitIntoLines(input)
	for idx, line := range lines {
		fmt.Printf("Line %d: %q\n", idx+1, line)
	}
}
