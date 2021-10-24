package main

import (
	"fmt"
	"sync"
)

func processAndGather(in <-chan int, processWG func(int) int,num int) []int {
	out := make(chan int, num) //"chan int" depend on return of "processWG"
	var wg sync.WaitGroup
	wg.Add(num)
	for i:=0; i<num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processWG(v)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []int
	for v:= range out {
		result = append(result,v)
	}
	return result
}

func processWg(val int) int {
	return val
}

func main() {
	inChan := make(chan int)

	go func() {
		for i:=0; i<6; i++ {
			inChan <- i
		}
		close(inChan)
	}()

	result := processAndGather(inChan,processWg,10)

	fmt.Println(result)
}
