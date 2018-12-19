package dummyservice

import (
	"math/rand"
)

// Worker represents a dummy service that can be pinged
type Worker struct{}

// Ping will give the status of this dummy worker. It will
// fail randomly just for demo purposes
func (s *Worker) Ping() bool {
	r := rand.Int31() % 10
	return r > 3
}

// New creates a dummy worker
func New() *Worker {
	return &Worker{}
}
