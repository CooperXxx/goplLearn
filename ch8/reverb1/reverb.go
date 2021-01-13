package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for  {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn,input.Text(),time.Second*1)
	}
	conn.Close()
}

func echo(conn net.Conn, text string, duration time.Duration) {
	fmt.Fprintln(conn,strings.ToUpper(text))
	time.Sleep(duration)
	fmt.Fprintln(conn,text)
	time.Sleep(duration)
	fmt.Fprintln(conn,strings.ToLower(text))
	time.Sleep(duration)
}

