package main

// Concurrency: Multiple process that execute indepently
// Example send 4 emails, one after the other, but do not wait for responses before moving to next email
// NOT sending the emails at the exact same time (NOT in parallel)
// Quote: "Concurrency is the composition of independently executing processes, while parallelism is the
// simultaneous execution of (possibly related) computations. Concurrency is about dealing with lots 
// of things at once. Parallelism is about doing lots of things at once.

// Multiple threads are scheduled to run on Cores (concurrency)
// Multiple core provide parallelism
// But Go does not use Threads for concurrency

// Go uses goroutines which are layered on top off threads
// Scheduled and managed by the go runtime (simpler)
// Lightweight, KBs rather than MBs
// Less thread switching. GoRoutine switching can happen on same thread. Less overhead, greater efficiency
// Faster start up times and can safely communicate across go routines

// Gos concurrency model: Actor model
// Communicating Sequential Processes (CSP)
// Actors pass messages between each other (each go routine) via channels

// Channels are like unix pipes or real world pipes
// One go routine puts message on pipe and one takes it off

import (
	"fmt"
	"time"
	"sync"
	"runtime"
)

func main() {

	// ****** This one line makes the program parallel ****** 
	// Remove this to just have concurrency without parallelism
	// runtime.GOMAXPROCS(2)

	// Below all works on a single thread with multiple go routines

	// Do not end function until everything in the wait group has reported that they are done
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // one for each function

	// anonymous self executing function
	// The go keyword make it a go routine which means that it will not block
	// 	execution will move onto the next function
	go func() {
		defer waitGroup.Done()
		
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGroup.Done()

		fmt.Println("World")
	}()

	// Do not end function until everything in the wait group has reported that they are done
	waitGroup.Wait()
}







