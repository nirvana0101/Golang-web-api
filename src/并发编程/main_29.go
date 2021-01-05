package main

import "fmt"

/**
select关键字的使用: 写、读，写、读，写、读……
*/
func calculate(ch chan<- int, ch_ <-chan int) {
	x, y := 1, 1
	for {
		select {
		case ch <- x: //1写
			x, y = y, x+y
		case num := <-ch_: //over
			fmt.Println(num)
			return
		}
	}
}
func main() {
	ch := make(chan int)
	ch_ := make(chan int)
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch //2读
			fmt.Println(num)
		}
		ch_ <- -1 //
	}()
	calculate(ch, ch_)
}
