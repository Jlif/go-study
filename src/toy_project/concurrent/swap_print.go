package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var dogCounter uint64
var catCounter uint64
var fishCounter uint64

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	dogCh := make(chan struct{})
	catCh := make(chan struct{})
	fishCh := make(chan struct{})

	go dog(&wg, dogCh, catCh)
	go cat(&wg, catCh, fishCh)
	go fish(&wg, fishCh, dogCh)

	dogCh <- struct{}{}

	wg.Wait()
}

func dog(wg *sync.WaitGroup, in chan struct{}, out chan struct{}) {
	for dogCounter < 10 {
		<-in
		fmt.Println("dog")
		atomic.AddUint64(&dogCounter, 1)
		out <- struct{}{}
	}
	close(in)
	wg.Done()
}

func cat(wg *sync.WaitGroup, in chan struct{}, out chan struct{}) {
	for catCounter < 10 {
		<-in
		fmt.Println("cat")
		atomic.AddUint64(&catCounter, 1)
		out <- struct{}{}
	}
	close(in)
	wg.Done()
}

func fish(wg *sync.WaitGroup, in chan struct{}, out chan struct{}) {
	for fishCounter < 10 {
		<-in
		fmt.Println("fish")
		atomic.AddUint64(&fishCounter, 1)
		if fishCounter != 10 {
			out <- struct{}{}
		}
	}
	close(in)
	wg.Done()
}
