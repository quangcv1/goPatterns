package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func timeLimit(f func() (int,error)) (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = f()
		close(done)
	}()

	select {
	case <-done:
		return result, nil
	case <-time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}
func doSomeWork() (int, error) {
	time.Sleep(1*time.Second)
	return 1,nil
}
func main() {
	val, err := timeLimit(doSomeWork)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}
