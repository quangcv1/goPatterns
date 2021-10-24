package main

import (
	"errors"
	"net/http"
	"time"
)
/**
Declare Backpressure pattern
 */
type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge  {
	ch := make(chan struct{}, limit)
	for i:=0; i< limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

func (pg *PressureGauge) Process(f func()) error  {
	select {
	case <- pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}
/**
Implement Backpressure
 */
func doThingThatShouldBeLimited() string {
	time.Sleep(5*time.Second)
	return "done"
}

func main() {
	pg := New(2)
	http.HandleFunc("/request", func(writer http.ResponseWriter, request *http.Request) {
		err := pg.Process(func() {
			writer.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			writer.WriteHeader(http.StatusTooManyRequests)
			writer.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
