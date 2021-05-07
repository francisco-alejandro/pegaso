package pegaso

import "sync"

// Bus handles susbcribers
type Bus struct {
	subscribers []chan interface{}
	m           sync.Mutex
}

// Subscribe adds access to a channel to the bus to receive messages
func (b *Bus) Subscribe(c chan interface{}) {
	b.m.Lock()

	defer b.m.Unlock()

	b.subscribers = append(b.subscribers, c)
}

// Publish a payload into the channel, each subscriber receives the payload
func (b *Bus) Publish(e interface{}) {
	go func() {
		b.m.Lock()
		defer b.m.Unlock()
		for _, s := range b.subscribers {
			s <- e
		}
	}()
}
