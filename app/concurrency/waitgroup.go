package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

//WaitGroup funcion for testing
func WaitGroup() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)
	fmt.Println("Start Goroutines")
	// Launch functions as goroutines
	go addTable()
	go multiTable()
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

func addTable() {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Addition Table for:", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d+%d=%d\t", i, j, i+j)
		}
		fmt.Println()
	}
}
func multiTable() {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Multiplication Table for:", i)
		for j := 1; j <= 10; j++ {
			//res = i + j
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
