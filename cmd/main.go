package main

import (
	"fmt"
	"github.com/mickeyyawn/dscribe"
)

func main() {

	for i := 0; i < 1000; i++ {
		s := dscribe.Generate()

		fmt.Print(s)
		fmt.Println()
	}

}
