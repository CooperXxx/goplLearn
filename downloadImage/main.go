package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var sema = make(chan struct{}, 4)

func downloadImage(url string) {
	defer func() { <-sema }() // release token
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get %s err!", url)
	}
	defer resp.Body.Close()
	names := strings.Split(url, "/")
	out, err := os.Create("e:" + names[len(names)-2] + "/" + names[len(names)-1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Creat pic file err!")
	}
	pic, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Save pic err!")
	}
	io.Copy(out, bytes.NewReader(pic))
}
func main() {
	urlModel := "http://121.248.63.139/photo/2020/"
	var url string
	for i := 4490; i < 6500; i++ {
		url = urlModel + strconv.Itoa(200000+i) + ".jpg"
		sema <- struct{}{} // acquire token
		//fmt.Println(url)
		go downloadImage(url)
	}

}
