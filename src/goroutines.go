package main

import (
	"fmt"
	"runtime"
	"github.com/czawadka/sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	workerCount := 100

	channel := make(chan int, 1)
	latch := sync.NewLatch()

	for i := 0; i < workerCount; i++ {
		go func(workerId int) {
			fmt.Printf("Worker %d inited\n", workerId)
			latch.Await()
			fmt.Printf("Worker %d started\n", workerId)
			channel <- workerId
			fmt.Printf("Worker %d stopped\n", workerId)
		}(i)
	}
	// unleash all workers
	fmt.Printf("Unleash %d routines at once\n", workerCount)
	latch.Release()

	for i := 0; i < workerCount; i++ {
		workerId := <- channel
		fmt.Printf("Worker %d reported\n", workerId)
	}

}
