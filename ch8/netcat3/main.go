package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer dial.Close()
	done := make(chan struct{})
	go func() {
		//不能使用mustCopy，因为must copy代码只会复制内容。不能处理eof或连接终中断
		//也不能这么说 因为mustcopy会抓到错误，所以不执行了 也就没有打印done和执行后续代码
		io.Copy(os.Stdout,dial)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(dial,os.Stdin)
	dial.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
