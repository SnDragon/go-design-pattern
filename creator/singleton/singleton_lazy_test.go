package singleton

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

type getInstanceFn func() *lazySingleton

func parallelGetInstance(t *testing.T, fn getInstanceFn) {
	size := 100
	clients := make([]*lazySingleton, size)
	wg := sync.WaitGroup{}
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(i int) {
			defer wg.Done()
			clients[i] = fn()
		}(i)
	}
	wg.Wait()
	c := clients[0]
	for i := 1; i < size; i++ {
		assert.Equal(t, c, clients[i])
	}
}

func TestGetLazySingletonInstanceWithDCL(t *testing.T) {
	parallelGetInstance(t, GetLazySingletonInstanceWithDCL)
}

func TestGetLazySingletonInstanceWithOnce(t *testing.T) {
	parallelGetInstance(t, GetLazySingletonInstanceWithOnce)
}
