package main

import (
	"fmt"
	"time"
)

var sharedCounter int

func incrementCounter() {
	for i := 0; i < 1000; i++ {
		sharedCounter++
	}
}

func decrementCounter() {
	for i := 0; i < 1000; i++ {
		sharedCounter--
	}
}

func main() {
	go incrementCounter()
	go decrementCounter()

	// Wait for goroutines to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Final value of sharedCounter:", sharedCounter)
}

/*
Race Condition Explanation:

The variable 'sharedCounter' is accessed and modified by two goroutines: 'incrementCounter' and 'decrementCounter'.
Both goroutines perform read-modify-write operations on 'sharedCounter' concurrently without any synchronization.
This can lead to a race condition where the final value of 'sharedCounter' is unpredictable and may not be zero,
because the operations (++ and --) are not atomic. The goroutines may interleave in such a way that increments and
decrements overwrite each other's changes, resulting in lost updates.
*/
