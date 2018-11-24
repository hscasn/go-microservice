package client1

import (
	"math/rand"
)

// Client1 is
type Client1 struct{}

// Ping is
func (s *Client1) Ping() bool {
	r := rand.Int31() % 10
	return r > 3
}

// Create is
func Create() *Client1 {
	return &Client1{}
}
