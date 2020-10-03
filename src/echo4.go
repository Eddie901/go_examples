// Echo4 prints its comman-line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trainiling newline")
var sep = flag.String("s", " ", "separator")
var c = flag.Int("c", 1, "times")

func main() {
	flag.Parse()
	for t := 0; t < *c; t += 1 {
		fmt.Print(strings.Join(flag.Args(), *sep))
		if !*n {
			fmt.Println()
		}
	}
}
