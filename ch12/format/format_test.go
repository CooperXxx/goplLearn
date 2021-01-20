package format_test

import (
	"fmt"
	"goplLearn/ch12/format"
	"testing"
)

func Test(t *testing.T) {
	//string
	fmt.Println(format.Any("hello"))
	//struct
	fmt.Println(format.Any(struct {
		s string
	}{"hello"}))
	//point
	fmt.Println(format.Any(func() { fmt.Println("hello-") }))
	//array
	fmt.Println(format.Any([3]int{1, 2, 3}))
	//slice
	fmt.Println(format.Any([]int{1, 2, 3}))
	//map
	fmt.Println(format.Any(map[string]int{"sss": 1}))

}
