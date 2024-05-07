package main

import (
	"fmt"

	"github.com/joseraulperezrodriguez/golang-useful-impl/app/concurrency"
)

func main() {

	//concurrency.WaitGroup()
	//fmt.Println(concurrency.UnboundedChannels())
	//fmt.Println(concurrency.BoundedChannels())
	//concurrency.ClosedChannnel()
	concurrency.DoFibonacci()
	fmt.Println("reached end")
}
