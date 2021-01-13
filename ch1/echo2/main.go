package main

import (
	"fmt"
	"os"
)

func main() {
	s,temp:="",""
	for _,arg:=range os.Args[1:] {
		s += temp + arg
		temp=" "
	}
	fmt.Println(s)
}
