package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func proccess(channel chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
	wg.Done()
}

func processCarRacing(channel chan int, wg *sync.WaitGroup) {
	min := 1
	max := 20
	for i := 0; i < 10; i++ {
		randomNumber := rand.Intn(max-min) + min
		channel <- randomNumber
	}
	wg.Done()
}

func worker(channel chan int, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for {
		fmt.Println("Received on channel", <-channel, " on workerID ", workerID)
	}
}

func workerCarRacing(channel chan int, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for {
		fmt.Println("Competitor ", workerID, " runs ", <-channel, " KM")
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

func loopCarRacing() {
	channel := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go processCarRacing(channel, &wg)
	go workerCarRacing(channel, &wg, 1)
	go workerCarRacing(channel, &wg, 2)
	wg.Wait()
}

func main() {
	loopCarRacing()

}
