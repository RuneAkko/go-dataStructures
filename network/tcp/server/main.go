package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic %s\n", p)
		}
	}()

	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(errors.New("can not Listen to 9999 for tcp"))
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		fmt.Println("shut down this tcp conn.")
		conn.Close()
	}()
	fmt.Println("new conn: ", conn.RemoteAddr())
	res := bytes.NewBuffer(nil)
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[0:])
		res.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				log.Println("read conn err: ", err)
			}
		}
		fmt.Println("receive: ", res)
		res.Reset()
	}
}
