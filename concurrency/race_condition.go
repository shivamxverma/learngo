package main

import "sync"

type Post struct {
	views int
	mu sync.Mutex
}


func (p* Post) inc(wg *sync.WaitGroup){
	defer func() {
		wg.Done()
		defer p.mu.Unlock()
	}()

	p.mu.Lock()
	
	p.views++
}

func runInc(){
	var wg sync.WaitGroup

	myPost := Post{views: 0}

	for i:=0 ; i < 100 ; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
}