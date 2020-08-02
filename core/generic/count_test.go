package generic

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCount(t *testing.T) {
	data := []interface{} {
		1, 2, 3, "Hello", "Hi", list.New(), '1', 2, 3,
	}

	counter := NewCount()

	// Add elements
	for _, v := range data {
		counter.Add(v)
	}

	res, err := counter.Result()
	ld := uint64(len(data))

	assert.Nil(t, err)
	assert.Equal(t, ld, res)

	// Remove elements
	counter.Remove(1)
	res, _ = counter.Result()
	assert.Equal(t, ld - 1, res)

	counter.Remove(1)
	res, _ = counter.Result()
	assert.Equal(t, ld - 1, res)
}
