package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/**
写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。
*/

func inter9() {
	wg := sync.WaitGroup{}
	nums := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			nums <- rand.Intn(10)
		}
		close(nums)
	}()

	go func() {
		defer wg.Done()
		for ele := range nums {
			fmt.Println(ele)
		}
	}()
	wg.Wait()
}
