package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numPhilosophers = 5
	numMeals        = 3
	maxEating       = 2
)

type Chopstick struct {
	sync.Mutex
	id int
}

type Philosopher struct {
	id              int
	leftChopstick   *Chopstick
	rightChopstick  *Chopstick
	hostRequestChan chan int
	hostReleaseChan chan int
}

func philosopherRoutine(p Philosopher, wg *sync.WaitGroup) {
	defer wg.Done()
	for meal := 0; meal < numMeals; meal++ {
		// Request permission from host
		p.hostRequestChan <- p.id

		// Pick up chopsticks in random order
		first, second := p.leftChopstick, p.rightChopstick
		if rand.Intn(2) == 0 {
			first, second = second, first
		}
		first.Lock()
		second.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(100)))
		fmt.Printf("finishing eating %d\n", p.id)

		second.Unlock()
		first.Unlock()

		// Notify host done eating
		p.hostReleaseChan <- p.id

		// Think for a bit
		time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(100)))
	}
}

func hostRoutine(requestChan, releaseChan chan int) {
	eating := 0
	queue := []int{}
	for {
		select {
		case id := <-releaseChan:
			eating--
			// Optionally, print host state for debugging
			_ = id
		case id := <-requestChan:
			if eating < maxEating {
				eating++
			} else {
				// Wait until a slot is free
				queue = append(queue, id)
				continue
			}
		}
		// Serve waiting philosophers if possible
		for eating < maxEating && len(queue) > 0 {
			eating++
			queue = queue[1:]
		}
		if eating == 0 && len(queue) == 0 {
			// Exit condition: all philosophers done
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = &Chopstick{id: i}
	}

	hostRequestChan := make(chan int, numPhilosophers)
	hostReleaseChan := make(chan int, numPhilosophers)

	var wg sync.WaitGroup
	wg.Add(numPhilosophers)

	// Start host goroutine
	go func() {
		activeEaters := 0
		waiting := []int{}
		for finished := 0; finished < numPhilosophers*numMeals; {
			select {
			case <-hostReleaseChan:
				activeEaters--
				finished++
			case id := <-hostRequestChan:
				if activeEaters < maxEating {
					activeEaters++
				} else {
					waiting = append(waiting, id)
					continue
				}
			}
			for activeEaters < maxEating && len(waiting) > 0 {
				activeEaters++
				waiting = waiting[1:]
			}
		}
	}()

	// Start philosopher goroutines
	for i := 0; i < numPhilosophers; i++ {
		p := Philosopher{
			id:              i + 1,
			leftChopstick:   chopsticks[i],
			rightChopstick:  chopsticks[(i+1)%numPhilosophers],
			hostRequestChan: hostRequestChan,
			hostReleaseChan: hostReleaseChan,
		}
		go philosopherRoutine(p, &wg)
	}

	wg.Wait()
}
