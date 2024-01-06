package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	user := fetchUser()
	fmt.Println("user: ", user)

	// ? setup channel an wait group (adn capacities)
	respChan := make(chan any, 2)
	waitGroup := &sync.WaitGroup{}

	waitGroup.Add(2)

	go fetchUserLikes(user, respChan, waitGroup)
	// fmt.Println("user likes:", userLikes)

	go fetchUserMatch(user, respChan, waitGroup)
	// fmt.Println("user match: ", userMatch)

	waitGroup.Wait() // ? block al proceso hasta se hayan 2 sync.Done()
	close(respChan)  // ? close channel

	for resp := range respChan {
		fmt.Println("resp: ", resp)
	}

	fmt.Println("total time: ", time.Since(startTime))
}

// ? the following functions mimic HTTP requests
// * channels are used sas the concurrency mechanism
// * standard sync coordinates the concurrency

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "BOB"
}

func fetchUserLikes(username string, respChan chan any, waitGroup *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	respChan <- 11

	// return 11

	waitGroup.Done()
}

func fetchUserMatch(username string, respChan chan any, waitGroup *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	respChan <- "ANNA"

	// return "ANNA"

	waitGroup.Done()
}
