package channelizer

import (
	"errors"
	"reflect"
)

type ChannelType uint8

const (
	ChannelTypeBytes  ChannelType = 1
	ChannelTypeInt    ChannelType = 2
	ChannelTypeString ChannelType = 3
)

type Channelizer struct {
	chanRoute    map[string]ChannelType
	chanRegistry map[string]any
}

func New() *Channelizer {
	return &Channelizer{
		chanRoute:    make(map[string]ChannelType),
		chanRegistry: make(map[string]any),
	}
}

func (c *Channelizer) Add(key string, channel any) error {
	switch channel.(type) {
	case chan []byte:
		c.chanRoute[key] = ChannelTypeBytes
	case chan int:
		c.chanRoute[key] = ChannelTypeInt
	case chan string:
		c.chanRoute[key] = ChannelTypeString
	default:
		return errors.New("unsupported data type")
	}
	c.chanRegistry[key] = channel
	return nil
}

func (c *Channelizer) Send(key string, data any) error {
	if _, isExist := c.chanRoute[key]; !isExist {
		return errors.New("key not found")
	}
	if !checkType(c.chanRegistry[key], data) {
		return errors.New("data type does not match")
	}
	switch c.chanRoute[key] {
	case ChannelTypeBytes:
		channel := c.chanRegistry[key].(chan []byte)
		channel <- data.([]byte)
	case ChannelTypeInt:
		channel := c.chanRegistry[key].(chan int)
		channel <- data.(int)
	case ChannelTypeString:
		channel := c.chanRegistry[key].(chan string)
		channel <- data.(string)
	}
	return nil
}

func checkType(channel any, data any) bool {
	return reflect.TypeOf(channel).Elem() == reflect.TypeOf(data)
}
