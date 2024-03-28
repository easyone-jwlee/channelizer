package main

import (
	"fmt"
	"time"

	"github.com/easyone-jwlee/channelizer"
)

func main() {
	chz := channelizer.New()

	channel1 := make(chan []byte)
	channel2 := make(chan int)
	channel3 := make(chan string)

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

	ticker1s := time.NewTicker(1 * time.Second)
	ticker2s := time.NewTicker(2 * time.Second)
	ticker3s := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker1s.C:
			if err := chz.Send("one", []byte("one")); err != nil {
				fmt.Printf("failed to send channel1. Error: %v\n", err)
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
