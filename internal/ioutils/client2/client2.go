package client2

import (
	"math/rand"
)

// Client2 is
type Client2 struct{}

// Ping is
func (s *Client2) Ping() bool {
	r := rand.Int31() % 10
	return r > 3
}

// Create is
func Create() *Client2 {
	return &Client2{}
}
