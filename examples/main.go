package main

import (
	"fmt"
	"time"

	"github.com/easyone-jwlee/channelizer"
)

func main() {
	chz := channelizer.New()

	channel1 := make(chan []byte)
	channel2 := make(chan []byte)

	chz.Add("one", channel1)
	chz.Add("two", channel2)

	go func() {
		for {
			select {
			case data := <-channel1:
				fmt.Printf("get data via channel1: %v\n", string(data))
			}
		}
	}()

	go func() {
		for {
			select {
			case data := <-channel2:
				fmt.Printf("get data via channel2: %v\n", string(data))
			}
		}
	}()

	ticker1s := time.NewTicker(1 * time.Second)
	ticker2s := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker1s.C:
			chz.Send("one", []byte("one"))
		case <-ticker2s.C:
			chz.Send("two", []byte("twotwo"))
		}
	}
}
