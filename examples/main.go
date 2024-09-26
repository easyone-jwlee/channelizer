package main

import (
	"fmt"
	"time"

	"github.com/easyone-jwlee/channelizer"
)

func main() {
	chz := channelizer.New()

	channel1 := make(chan []byte, 100)
	channel2 := make(chan int, 10)
	channel3 := make(chan string, 10)

	if err := chz.Add("one", channel1); err != nil {
		fmt.Printf("failed to add channel1. Error: %v\n", err)
		return
	}
	if err := chz.Add("two", channel2); err != nil {
		fmt.Printf("failed to add channel2. Error: %v\n", err)
		return
	}
	if err := chz.Add("three", channel3); err != nil {
		fmt.Printf("failed to add channel3. Error: %v\n", err)
		return
	}

	ticker1s := time.NewTicker(1 * time.Second)
	ticker2s := time.NewTicker(2 * time.Second)
	ticker3s := time.NewTicker(3 * time.Second)
	ticker10s := time.NewTicker(10 * time.Second)

	go func() {
		countOne := 0
		for {
			select {
			case <-channel1:
				countOne++
			case <-ticker10s.C:
				fmt.Printf("get data via channel1: %v\n", countOne)
				countOne = 0
			}
		}
	}()

	go func() {
		for {
			select {
			case data := <-channel2:
				fmt.Printf("get data via channel2: %v\n", data)
			}
		}
	}()

	go func() {
		for {
			select {
			case data := <-channel3:
				fmt.Printf("get data via channel3: %v\n", data)
			}
		}
	}()

	for {
		select {
		case <-ticker2s.C:
			if err := chz.MonitorChannelBuffer("one"); err != nil {
				fmt.Printf("failed to monitor buffer of channel1. Error: %v\n", err)
			}
			if err := chz.MonitorChannelBuffer("two"); err != nil {
				fmt.Printf("failed to monitor buffer of channel1. Error: %v\n", err)
			}
			if err := chz.MonitorChannelBuffer("three"); err != nil {
				fmt.Printf("failed to monitor buffer of channel1. Error: %v\n", err)
			}
		case <-ticker1s.C:
			for i := 0; i < 50000; i++ {
				if err := chz.Send("one", []byte("one")); err != nil {
					fmt.Printf("failed to send channel1. Error: %v\n", err)
				}
			}
		case <-ticker2s.C:
			if err := chz.Send("two", 2); err != nil {
				fmt.Printf("failed to send channel2. Error: %v\n", err)
			}
		case <-ticker3s.C:
			if err := chz.Send("three", "three"); err != nil {
				fmt.Printf("failed to send channel3. Error: %v\n", err)
			}
		}
	}
}
