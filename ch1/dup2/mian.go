package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	count := make(map[string]int)
	file := os.Args[1:]
	if len(file) ==0  {
		countWords(os.Stdin,count)
	} else {
		for _,arg := range file{
			f,err := os.Open(arg)
			if err!=nil {
				fmt.Printf("dup2:%v\n",err)
				continue
			}
			countWords(f,count)
			f.Close()
		}
	}
	for words,n := range count{
		if n>1 {
			fmt.Printf("%s:%dtimes\n",words,n)
		}
	}
}

func countWords(f *os.File, count map[string]int) {
	input:=bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
}
