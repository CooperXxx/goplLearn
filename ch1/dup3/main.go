package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(reflect.TypeOf(reflect.TypeOf("aa")))
	fmt.Println(reflect.TypeOf(reflect.TypeOf("aa").Kind()))
	fmt.Println(reflect.TypeOf(reflect.TypeOf("aa")).Kind())
	fmt.Println(reflect.TypeOf(reflect.TypeOf("aa").Kind()).Kind())
	fmt.Println(nil)
	count := make(map[string]int)
	for _,filename := range os.Args[1:]{
		data , err := ioutil.ReadFile(filename)
		if err!= nil  {
			fmt.Printf("dip3:%v",err)
			continue
		}
		for _,line := range strings.Split(string(data),"\n"){
			count[line]++
		}

	}

	for words,n := range count{
		if n>1 {

			fmt.Printf("%s\n%d\n",words,n)
		}
	}
	

}
