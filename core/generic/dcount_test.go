package generic

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDCount(t *testing.T) {
	data := []interface{} {
		1, 2, 3, "Hello", "Hi", list.New(), '1', 2, 3,
	}

	counter := NewDCount()

	// Add elements
	for _, v := range data {
		counter.Add(v)
	}

	res, err := counter.Result()
	ld := uint64(len(data)) - 2

	assert.Nil(t, err)
	assert.Equal(t, ld, res)

	// Remove elements
	counter.Remove(2)
	res, _ = counter.Result()
	assert.Equal(t, ld, res)

	counter.Remove(1)
	res, _ = counter.Result()
	assert.Equal(t, ld - 1, res)

	counter.Remove(1)
	res, _ = counter.Result()
	assert.Equal(t, ld - 1, res)
}
