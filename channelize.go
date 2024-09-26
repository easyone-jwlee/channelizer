package channelizer

import (
	"fmt"
	"reflect"
)

type ChannelType uint8

const (
	ChannelTypeBytes ChannelType = iota + 1
	ChannelTypeInt
	ChannelTypeString
)

type ChannelData struct {
	ChannelType ChannelType
	Channel     any
}

type Channelizer struct {
	chanRegistry map[string]ChannelData
}

func New() *Channelizer {
	return &Channelizer{
		chanRegistry: make(map[string]ChannelData),
	}
}

func (c *Channelizer) Add(key string, channel any) error {
	var channelType ChannelType
	switch channel.(type) {
	case chan []byte:
		channelType = ChannelTypeBytes
	case chan int:
		channelType = ChannelTypeInt
	case chan string:
		channelType = ChannelTypeString
	default:
		return fmt.Errorf("unsupported data type. data_type: %v", reflect.TypeOf(channel))
	}
	c.chanRegistry[key] = ChannelData{
		ChannelType: channelType,
		Channel:     channel,
	}
	return nil
}

func (c *Channelizer) Send(key string, data any) error {
	if _, isExist := c.chanRegistry[key]; !isExist {
		return fmt.Errorf("key not found")
	}
	if err := c.checkType(key, data); err != nil {
		return fmt.Errorf("failed to check data type. error: %v", err)
	}
	switch c.chanRegistry[key].ChannelType {
	case ChannelTypeBytes:
		channel := c.chanRegistry[key].Channel.(chan []byte)
		channel <- data.([]byte)
	case ChannelTypeInt:
		channel := c.chanRegistry[key].Channel.(chan int)
		channel <- data.(int)
	case ChannelTypeString:
		channel := c.chanRegistry[key].Channel.(chan string)
		channel <- data.(string)
	}
	return nil
}

func (c *Channelizer) checkType(key string, data any) error {
	switch c.chanRegistry[key].ChannelType {
	case ChannelTypeBytes:
		if _, ok := data.([]byte); !ok {
			return fmt.Errorf("data type mismatch: expected []byte")
		}
	case ChannelTypeInt:
		if _, ok := data.(int); !ok {
			return fmt.Errorf("data type mismatch: expected int")
		}
	case ChannelTypeString:
		if _, ok := data.(string); !ok {
			return fmt.Errorf("data type mismatch: expected string")
		}
	}
	return nil
}

func (c *Channelizer) MonitorChannelBuffer(key string) error {
	if _, isExist := c.chanRegistry[key]; !isExist {
		return fmt.Errorf("key not found")
	}
	switch c.chanRegistry[key].ChannelType {
	case ChannelTypeBytes:
		channel := c.chanRegistry[key].Channel.(chan []byte)
		fmt.Printf("Channel buffer usage: %d/%d\n", len(channel), cap(channel))
	case ChannelTypeInt:
		channel := c.chanRegistry[key].Channel.(chan int)
		fmt.Printf("Channel buffer usage: %d/%d\n", len(channel), cap(channel))
	case ChannelTypeString:
		channel := c.chanRegistry[key].Channel.(chan string)
		fmt.Printf("Channel buffer usage: %d/%d\n", len(channel), cap(channel))
	}
	return nil
}
