package main

import (
	"fmt"
	"sync"
)

var (
	catCh  = make(chan struct{})
	fishCh = make(chan struct{})
	dogCh  = make(chan struct{})
)
var wg sync.WaitGroup

func cat() {
	for range catCh {
		fmt.Print("cat ")
		dogCh <- struct{}{}
	}
}

func dog() {
	for range dogCh {
		fmt.Print("dog ")
		fishCh <- struct{}{}
	}
}

func fish() {
	defer wg.Done()
	i := 0
	for range fishCh {
		i++
		fmt.Print("fish ")
		fmt.Println(i)
		if i == 100 {
			break
		}

		catCh <- struct{}{}
	}

}

func main() {
	wg.Add(1)
	go dog()
	go cat()
	go fish()
	catCh <- struct{}{}
	wg.Wait()
}
