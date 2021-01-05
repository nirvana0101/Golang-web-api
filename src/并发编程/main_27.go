package main

import "fmt"

/**
关闭通道
*/
func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}

/**
关闭通道
*/
func main_() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		if i, ok := <-ch; ok == true {
			fmt.Println(i)
		} else {
			break
		}
	}
}
