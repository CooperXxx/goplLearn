package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input:= bufio.NewScanner(os.Stdin)
	count := make(map[string]int)
	for input.Scan() {
		count[input.Text()]++
	}
	for line,n := range count{
		if n>1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
