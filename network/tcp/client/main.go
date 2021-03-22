package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	// data := []byte("[123456789123456789123456789123456789]")
	// conn, err := net.DialTimeout("tcp", "localhost:9999", time.Second*30)
	// if err != nil {
	// 	log.Printf("conn failed: %s\n", err)
	// }
	// for j := 0; j < 99; j++ {
	// 	_, err = conn.Write(data)
	// 	if err != nil {
	// 		break
	// 	}
	// }
	conn, err := net.DialTimeout("tcp", "localhost:9999", time.Second*30)
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 10; i++ {
		go func(i int) {
			str := "[" + strings.Repeat(fmt.Sprintf("%d", i), 100) + "]"

			if err != nil {
				log.Printf("conn failed: %s\n", err)
			}
			for j := 0; j < 99; j++ {
				_, err = conn.Write([]byte(str))
				if err != nil {
					break
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
