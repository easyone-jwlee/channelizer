# Channelizer

Channelizer is a Go package designed to simplify and enhance the use of channels in applications where imports are frequent and channels are extensively used. The primary goal of Channelizer is to mitigate the complexity that can arise from such architectures, often indicated by high cyclomatic complexities (`import cycle`). This complexity arises from the intricate web of dependencies and the heavy use of channels across different parts of an application.

To achieve its goal, Channelizer introduces a novel approach by employing a map that uses strings as keys to register and reference channels. This mechanism allows for an organized and scalable way to manage communication channels within your application, enabling you to send data to specific channels identified by keys.

## Features

* **Simplified Channel Management:** Utilizes a string-keyed map to manage channels, making it easier to reference and send data across different parts of an application.
* **Byte Slice Data Handling:** Initially supports sending data in the form of []byte, catering to a wide range of applications that require binary data communication.
* **Future Enhancements:** Plans to support all data types, enabling the registration and transmission of any Go data type through the managed channels.

## Getting Started

This section provides a quick guide on how to integrate Channelizer into your Go application.

### Installation

To install Channelizer, use the `go get` command:

```bash
go get -u github.com/easyone-jwlee/channelizer
```

## Usage

Here's a simple example to get you started with Channelizer:

```go
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
```

This example demonstrates how to create a new Channelizer instance, register a channel with a unique key, and send and receive data using that key.

To test, run:

```bash
make test

```

## Future Directions

Channelizer aims to evolve into a comprehensive solution for managing inter-process communication in Go applications. Future versions will introduce the ability to handle all Go data types, further simplifying the process of sending and receiving data across channels. This flexibility will empower developers to build more complex and responsive applications with ease.

## Contributing

Contributions are welcome! If you have ideas on how to improve Channelizer or want to contribute code, please feel free to submit issues and pull requests on GitHub.

## License

Channelizer is released under the MIT License. See the LICENSE file for more details.

---

This README provides an overview of Channelize, emphasizing its simplicity and potential for future enhancements. Adjust the installation instructions and examples as necessary to match the actual implementation and repository location.