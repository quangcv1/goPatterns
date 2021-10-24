package main

import (
	"fmt"
)

/**
How to we want to terminate the goroutine while executing
With done pattern we use it to signal to exit goroutine so
=> we should use function to wrap done channel with close
then we can call it from outside to signal
=> return at least a function with wrapped "close(done)"
 */
/**
Note: with return a channel, make sure that channel should be closed
before we call it with "for range loop"
 */
func countTo(max int) (chan int, func())  {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
		close(ch)
	}
	go func() {
		for i:=0; i<max; i++ {
			select {
			case <-done: //when we close(done) then <-done will have value 0
				return //return to exit "go func()" routine
			case ch <- i:
			}
		}
		//close(ch)
	}()
	return ch, cancel //return to exit "go countTo func" routine
}
/**
checkChannel is used to check the channel is closed or not
why we need this function because:
1. We can't use "val,ok := <- ch" on main goroutine because
it will cause to "deadlock" since main goroutine will wait forever until have
another goroutine write
 */
func checkChannel(test *chan int) (int,bool) {
	var val int
	var ok bool
	val, ok = func(test1 *chan int) (int,bool){
		val, ok = <- *test1
		fmt.Printf("checkChannel: %v %v\n",val,ok)
		return val,ok
	}(test)
	fmt.Printf("out checkchannel: %v %v\n", val,ok)
	return val,ok
}

func main() {
	ch, cancel := countTo(10)
	//checkChan := make(chan int)
	for val := range ch {
		if val>5 {
			break
		}
		fmt.Println(val)
	}
	cancel()
	v,ok := <-ch
	fmt.Println(v,ok)

	//checkChan := make(chan int)
	//checkChan <- 1
	//fmt.Println(<-checkChan)

}
