package main

import (
	"flag"
	"fmt"
)

func main() {
	nFlag := flag.Int("n", 23, "start number")
	flag.Parse()

	n := *nFlag
	fmt.Println(n)

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		fmt.Println(n)
	}
}
