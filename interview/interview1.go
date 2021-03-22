package main

import (
	"fmt"
	"sync"
)

func interview1() {
	letter, num := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		i := 1
		for {
			select {
			case <-num:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	go func(doneFunc func()) {
		strIntStart := 'A'
		strIntEnd := 'Z'
		i := strIntStart
		for {
			select {
			case <-letter:
				if i > strIntEnd {
					doneFunc()
					return
				}
				fmt.Print(string(i), string(i+1))
				i += 2
				num <- true
				break
			default:
				break
			}
		}
	}(wg.Done)
	num <- true
	wg.Wait()
}
