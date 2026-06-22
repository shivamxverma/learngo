package main

import (
	"fmt"
	"sync"
	"time"
)

var urls = []string{
	"image.png",
	"image1.png",
}

func runWorker(jobsChan chan string, wg *sync.WaitGroup, resChan chan string) {
	defer wg.Done()

	for job := range jobsChan {
		time.Sleep(time.Millisecond * 50)
		fmt.Printf("Image processed: %s\n", job)
		resChan <- job
	}

	fmt.Println("Worker shutting down")
}

func StartWorker() {
	var wg sync.WaitGroup
	jobsChan := make(chan string, len(urls))
	resChan := make(chan string, len(urls))
	totalWorkers := 5

	start := time.Now()

	for i := 1 ; i <= totalWorkers ; i++ {
		wg.Add(1)
		go runWorker(jobsChan, &wg, resChan)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	for i := 0 ; i < len(urls) ; i++ {
		jobsChan <- urls[i]
	} 


	fmt.Println("Worker Finished", time.Since(start))
}


