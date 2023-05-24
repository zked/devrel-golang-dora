package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	waitForServer()

	for {
		var wg sync.WaitGroup
		numRequests := 100
		wg.Add(numRequests)
		for i := 0; i < numRequests; i++ {
			go func() {
				defer wg.Done()
				requestGames()
				requestGamesRandomId()
			}()
		}
		wg.Wait()
		time.Sleep(1 * time.Second)
	}
}

func requestGames() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("http://localhost:8081/games") // Replace with the appropriate endpoint URL
	if err != nil {
		log.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status: %s", resp.Status)
		return
	}

	log.Println("Request successful Games")
}

func requestGamesRandomId() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(4) + 1
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("http://localhost:8081/games/" + strconv.Itoa(randomNumber)) // Replace with the appropriate endpoint URL
	if err != nil {
		log.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status: %s", resp.Status)
		return
	}

	log.Println("Request successful GameID")
}

func waitForServer() {
	retryInterval := 1 * time.Second
	for {
		_, err := http.Get("http://localhost:8081/games") // Replace with appropriate health endpoint URL
		if err == nil {
			log.Println("Server is available")
			return
		}

		log.Printf("Server not available, retrying in %s...", retryInterval)
		time.Sleep(retryInterval)
	}
}
