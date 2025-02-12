package main

import "math/rand"

type Connection struct {
}

func (c *Connection) getItem(key string) Rankable {
	// Any implemenation is fine for this exercise.
	return Rankable{
		key:  key,
		rank: rand.Intn(100),
	}
}
