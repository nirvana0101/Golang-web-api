package main

import (
	"fmt"
	"time"
)

/**
有缓存的channel,先进先出，边存边取，缓存满了就会阻塞
*/
func main() {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //缓存满了就会阻塞
			fmt.Println("go:", i)
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-ch //没缓存就会阻塞
		fmt.Println("main:", num)
	}
}
