package main

import (
	"fmt"
	"time"
)

func processChannel(ch chan int) []int  {
	const conc = 10
	results := make(chan int,conc)
	for i:=0; i<conc; i++ {
		go func() {
			v := <- ch
			results <- processConc(v)
		}()
	}
	var out []int
	for i:=0;i<conc;i++ {
		out = append(out,<-results)
	}
	return out
}

func processConc(val int) int{
	return val
}

func main() {
	inChan := make(chan int)
	for i:=0; i<10;i++{
		go func(val int) {
			inChan<-val
		}(i)
	}
	time.Sleep(1*time.Second)
	out := processChannel(inChan)
	fmt.Println(out)
}
