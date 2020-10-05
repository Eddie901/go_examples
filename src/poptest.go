package main

import (
	"fmt"
	"os"
	"popcount"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseUint(arg, 0,64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "poptest: %v\n", err)
		}
		fmt.Printf("Popcount of %v %064b is %d\n",
			t, t, popcount.PopCount(t))
	}
}
