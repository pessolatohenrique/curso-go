package main

import (
	"curso-go/internal/entity"
	"fmt"
	"sync"
	"time"
)

func proccess(channel chan int, wg *sync.WaitGroup) {
	for i := 0; i < 100000; i++ {
		channel <- i
	}
	close(channel)
	wg.Done()
}

func worker(channel chan int, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for {
		fmt.Println("Received on channel", <-channel, " on workerID ", workerID)
	}
}

func loopConcurrency() {
	startTime := time.Now()

	channel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go proccess(channel, &wg)
	go worker(channel, &wg, 1)
	go worker(channel, &wg, 2)
	go worker(channel, &wg, 3)
	go worker(channel, &wg, 4)
	wg.Wait()

	duration := time.Since(startTime)
	fmt.Println("Duration of execution", duration)
}

func main() {
	// concurrency
	loopConcurrency()

	// order
	order, err := entity.NewOrder("123", 50, 10)
	if err == nil {
		order.CalculateFinalPrice()
		fmt.Println(order)
		fmt.Println(order.ID)
		fmt.Println(order.FinalPrice)
	} else {
		fmt.Println(err)
	}

}
