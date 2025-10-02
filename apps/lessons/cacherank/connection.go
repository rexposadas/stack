package main

import "math/rand"

type Connection struct {
}

func (c *Connection) getItem(key string) (Rankable, error) {
	// Any implementation is fine for this exercise.
	return Rankable{
		Key:  key,
		Rank: rand.Intn(100),
	}, nil
}
