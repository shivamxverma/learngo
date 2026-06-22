package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	value string
	Err error
}

func worker(url string, wg *sync.WaitGroup, resChan chan Result) {
	defer wg.Done()
	sum := 0

	for i := 0; i < 5_000_000_000; i++ {
		sum += i
	}
	resChan <- Result{
		value: url,
		Err:   nil,
	}
}

func main(){
	var wg sync.WaitGroup
	resChan := make(chan Result, 5)
	start := time.Now()

	numWorkers := 5
	wg.Add(numWorkers)

	// for i := range numWorkers {
	// 	go worker(fmt.Sprintf("image%d.png", i), &wg)
	// }
	go worker("image.png", &wg, resChan)
	go worker("image1.png", &wg, resChan)
	go worker("image2.png", &wg, resChan)
	go worker("image3.png", &wg, resChan)
	go worker("image4.png", &wg, resChan)
	// go worker("image5.png", &wg)
	// go worker("image6.png", &wg)
	// go worker("image7.png", &wg)
	// go worker("image8.png", &wg)

	wg.Wait()
	close(resChan)

	for res := range resChan {
		fmt.Printf("Received: %s\n", res.value)
	}
	fmt.Printf("It took %s time\n", time.Since(start))
}
