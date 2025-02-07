package pubsub

import "sync"

type PubSub struct {
	subscribers map[string][]chan string
	mux         sync.Mutex
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan string),
		mux:         sync.Mutex{},
	}
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
	ps.mux.Lock()
	defer ps.mux.Unlock()
	ch := make(chan string, 10)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return ch
}

func (ps *PubSub) Publish(topic, message string) {
	ps.mux.Lock()
	defer ps.mux.Unlock()
	for _, ch := range ps.subscribers[topic] {
		ch <- message
	}
}
