package main

import (
	"fmt"
	"testing"
)

func TestRangeOverClosedChannel(t *testing.T) {
	ch := make(chan int, 10)
	for i := 0; i != 10; i++ {
		ch <- i
	}
	close(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
