package main

import (
	"fmt"
	"time"
)

/**
Search Data voi string s with many call asynchronous and return if any one finish
Ex: return []string with
process func(string) []string
 */

func searchData(s string,searchers []func(string) []string) []string  {
	//step1: create done and result channel
	done := make(chan struct{})
	result := make(chan []string)
	for _,searcher := range searchers {
		go func(searcher func(string2 string) []string) {
			select {
			case result <- searcher(s):
			case <-done:
			}
		}(searcher)
	}
	r := <- result
	return r
}
/**
simulate 2 function with time sleep to see what the returned result
 */
func test(string) []string {
	time.Sleep(4*time.Second)
	return []string{"a","b"}
}

func test1(string) []string  {
	time.Sleep(2*time.Second)
	return []string{"c","d"}
}

func main() {
	 testArr := []func(string) []string{
		test,
		test1,
	}

	testData := searchData("a",testArr)

	fmt.Println(testData)
}
