package main

import (
	"fmt"
	"sync"
	"time"
)
/**
Done pattern: use for a single goroutine
WaitGroup pattern: wait on several goroutines
-- Add : increments the counter of goroutine to wait for
-- Done: decrements the counter and is called by a goroutine when it is finished
++Note: defer is called even if the goroutine panics
 */
func waitMultiGoroutines() {
	var wg sync.WaitGroup //this is variable is accessed by closure
	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(2*time.Second)
		dothing(1)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(2*time.Second)
		dothing(2)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(3*time.Second)
		dothing(3)
	}()
	wg.Wait()
}

func dothing(val int) {
	fmt.Println(val)
}

func main() {
	waitMultiGoroutines()
}
