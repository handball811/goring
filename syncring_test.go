package goring_test

import (
	"sync"
	"testing"

	. "github.com/handball811/goring"
	"github.com/stretchr/testify/assert"
)

var (
	N = 10000
)

func TestSyncPushAndPop(t *testing.T) {
	// setup
	target := NewSyncRing()

	// when
	var waitGroup sync.WaitGroup
	for i := 0; i < N; i++ {
		waitGroup.Add(1)
		go func() {
			target.Push("a")
			target.Pop()
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()

	// then
	assert.Equal(t, 0, target.Len())
}

func BenchmarkSyncPushAndPop(b *testing.B) {
	// setup
	target := NewSyncRing()
	var waitGroup sync.WaitGroup
	f := func() {
		target.Push("a")
		target.Pop()
		waitGroup.Done()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		waitGroup.Add(1)
		go f()
	}
	waitGroup.Wait()
}
