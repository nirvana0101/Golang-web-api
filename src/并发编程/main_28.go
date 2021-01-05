package main

/**
单向写
*/
func producer(w chan<- int) {
	for i := 0; i < 10; i++ {
		w <- i
	}
	close(w)
}

/**
单向读
*/
func consumer(r <-chan int) {
	for i := range r {
		println(i)
	}
}

/**
单向通道
*/

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
