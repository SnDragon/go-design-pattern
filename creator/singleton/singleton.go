package singleton

import (
	"fmt"
	"sync"
)

type client struct {
}

var once sync.Once
var c *client

func NewClient() *client {
	if c != nil {
		return c
	}
	once.Do(func() {
		fmt.Println("client init...")
		c = &client{}
	})
	return c
}
