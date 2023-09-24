package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	numUsers := 10
	numRequests := 100

	var wg sync.WaitGroup

	for i := 0; i < numUsers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stressTest(numRequests)
		}()
	}

	wg.Wait()
	fmt.Println("Stress testing completed.")
}

func stressTest(numRequests int) {
	targetURL := "https://your-shop-website.com"

	for i := 0; i < numRequests; i++ {
		startTime := time.Now()

		resp, err := http.Get(targetURL)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		defer resp.Body.Close()

		elapsedTime := time.Since(startTime)
		fmt.Printf("Request #%d - Status: %s, Elapsed Time: %v\n", i+1, resp.Status, elapsedTime)

	}
}
