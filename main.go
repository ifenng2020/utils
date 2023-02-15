package main

import (
	"sync"
	"time"
)

var timer = time.NewTimer(5 * time.Second)

func consumer(ch <-chan bool) bool {
	select {
	case d := <-ch:
		if !d {
			println("recv false")
			return true
		}
		println("recv true")
		return false
	case <-timer.C:
		println("time expired")
		return true
	}
}

func main() {
	c := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	// productor
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			c <- false
		}
		time.Sleep(time.Second)
		c <- true
		wg.Done()
	}()
	// consumer
	go func() {
		for {
			if b := consumer(c); !b {
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
}
