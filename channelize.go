package channelizer

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

func (c *Channelizer) Send(key string, data []byte) {
	c.chanRegistry[key] <- data
}
