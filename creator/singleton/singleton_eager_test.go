package singleton

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestGetEagerSingletonInstance(t *testing.T) {
	size := 100
	clients := make([]*eagerSingleton, size)
	wg := sync.WaitGroup{}
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(i int) {
			defer wg.Done()
			clients[i] = GetEagerSingletonInstance()
		}(i)
	}
	wg.Wait()
	c := clients[0]
	for i := 1; i < size; i++ {
		assert.Equal(t, c, clients[i])
	}
}
