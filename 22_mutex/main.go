package main

import (
	"fmt"
	"sync"
)

type Post struct {
	Views int
	mux   sync.Mutex
}

func (p *Post) IncrementViews(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mux.Unlock()
	}()
	p.mux.Lock()
	p.Views++
}

func main() {
	wg := sync.WaitGroup{}
	p := Post{Views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		p.IncrementViews(&wg)
	}
	wg.Wait()
	fmt.Println("Post views:", p.Views)
}
