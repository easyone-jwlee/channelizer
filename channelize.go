package channelizer

import (
	"errors"
)

type Channelizer struct {
	chanRegistry map[string]chan []byte
}

func New() *Channelizer {
	return &Channelizer{
		chanRegistry: make(map[string]chan []byte),
	}
}

func (c *Channelizer) Add(key string, channel chan []byte) {
	c.chanRegistry[key] = channel
}

func (c *Channelizer) Send(key string, data []byte) error {
	if _, isExist := c.chanRegistry[key]; !isExist {
		return errors.New("key not found")
	}
	c.chanRegistry[key] <- data
	return nil
}
