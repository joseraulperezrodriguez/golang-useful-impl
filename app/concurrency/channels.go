package concurrency

import (
	"fmt"
	"sync"
)

func addToChannel(channel chan int, val int) {
	channel <- val
}

func UnboundedChannels() int {
	var channel chan int = make(chan int)
	go addToChannel(channel, 2)
	go addToChannel(channel, 1)

	a, b := <-channel, <-channel

	return a + b
}

func BoundedChannels() int {
	var channel chan int = make(chan int, 1)
	go addToChannel(channel, 2)
	go addToChannel(channel, 1)

	a, b := <-channel, <-channel

	return a + b
}

func consume(input chan int, output chan int, wg *sync.WaitGroup) {
	ans := 0
	for {
		val, ok := <-input
		if !ok {
			break
		}
		ans += val
	}
	output <- ans
	wg.Done()
}

func feedback(input chan int, output chan int) {
	total := 0
	for {
		temp, ok := <-input
		if !ok {
			break
		}
		total += temp
	}
	output <- total
}

func ClosedChannnel() {
	twoMillions := 2000 * 1000
	routines := 20
	var input chan int = make(chan int, twoMillions)
	var output chan int = make(chan int, routines)
	for i := 0; i < twoMillions; i++ {
		input <- 1
	}
	close(input)
	consumerWaitGroup := &sync.WaitGroup{}
	consumerWaitGroup.Add(routines)
	for i := 0; i < routines; i++ {
		go consume(input, output, consumerWaitGroup)
	}

	aggr := make(chan int, 1)
	go feedback(output, aggr)
	consumerWaitGroup.Wait()
	close(output)

	ans := <-aggr
	fmt.Println(ans)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func DoFibonacci() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
