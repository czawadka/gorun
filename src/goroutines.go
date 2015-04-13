package main

import (
	"fmt"
	"github.com/czawadka/sync"
	"runtime"
	"time"
)

func main() {
	start(2)
}

func start(workerCount int) {
	runtime.GOMAXPROCS(8)
	fmt.Printf("GOMAXPROCS %d\n", runtime.GOMAXPROCS(0))

	channel := make(chan int, workerCount)
	closeChannel := func() {
		fmt.Printf("Closing bbb channel\n")
		close(channel)
	}
	defer closeChannel()

	latch := sync.NewCountDownLatch(1)

	for i := 0; i < workerCount; i++ {
		go func(workerId int) {
			fmt.Printf("Worker %d inited\n", workerId)
			latch.Await()
			fmt.Printf("Worker %d started\n", workerId)
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("Worker %d stopping\n", workerId)
			channel <- workerId
			fmt.Printf("Worker %d stopped\n", workerId)
		}(i)
	}
	// unleash all workers
	fmt.Printf("Unleash %d routines at once\n", workerCount)
	latch.CountDown()

	for i := 0; i < workerCount; i++ {
		workerId := <- channel
		fmt.Printf("Worker %d reported\n", workerId)
	}

	fmt.Printf("Done!\n")
}
