package goring_test

import (
	"testing"

	. "github.com/handball811/goring"
	"github.com/stretchr/testify/assert"
)

func TestCap(t *testing.T) {
	// setup
	target := NewRingWithSize(128)

	// when
	cap := target.Cap()

	//then
	assert.Equal(t, 128, cap)
}

func TestSetCapLarger(t *testing.T) {
	// setup
	target := NewRingWithSize(128)

	// when
	target.SetCap(256)
	cap := target.Cap()

	//then
	assert.Equal(t, 256, cap)
}

func TestSetCapSmaller(t *testing.T) {
	// setup
	target := NewRingWithSize(128)

	// when
	target.SetCap(64)
	cap := target.Cap()

	//then
	assert.Equal(t, 128, cap)
}

func TestPushAndLen(t *testing.T) {
	// setup
	target := NewRingWithSize(128)

	// when
	target.Push("a")

	//then
	assert.Equal(t, 1, target.Len())
}

func TestPushSliceAndLen(t *testing.T) {
	// setup
	target := NewRingWithSize(128)

	// when
	target.PushSlice([]interface{}{"a", "b", "c"})

	//then
	assert.Equal(t, 3, target.Len())
}

func TestPushAndPop(t *testing.T) {
	// setup
	target := NewRingWithSize(128)

	// when
	target.PushSlice([]interface{}{"a", "b", "c"})
	s, ok := target.Pop()

	//then
	assert.Equal(t, 2, target.Len())
	assert.Equal(t, "a", s.(string))
	assert.Equal(t, true, ok)
}

func TestPushAndPopSlice(t *testing.T) {
	// setup
	pops := make([]interface{}, 2)
	target := NewRingWithSize(128)

	// when
	target.PushSlice([]interface{}{"a", "b", "c"})
	size := target.PopSlice(pops)

	//then
	assert.Equal(t, 1, target.Len())
	assert.Equal(t, 2, size)
	assert.Equal(t, "a", pops[0].(string))
	assert.Equal(t, "b", pops[1].(string))
}

func TestPushAndPopSlice2(t *testing.T) {
	// setup
	pops := make([]interface{}, 4)
	target := NewRingWithSize(128)

	// when
	target.PushSlice([]interface{}{"a", "b", "c"})
	size := target.PopSlice(pops)

	//then
	assert.Equal(t, 0, target.Len())
	assert.Equal(t, 3, size)
	assert.Equal(t, "a", pops[0].(string))
	assert.Equal(t, "b", pops[1].(string))
	assert.Equal(t, "c", pops[2].(string))
}

func TestAutoGrowInPush(t *testing.T) {

	// setup
	pops := make([]interface{}, 4)
	target := NewRingWithSize(1)

	// when
	target.Push("a")
	target.Push("b")
	target.Push("c")
	target.Push("d")
	size := target.PopSlice(pops)

	//then
	assert.Equal(t, 0, target.Len())
	assert.Equal(t, 4, size)
	assert.Equal(t, "a", pops[0].(string))
	assert.Equal(t, "b", pops[1].(string))
	assert.Equal(t, "c", pops[2].(string))
	assert.Equal(t, "d", pops[3].(string))
}
