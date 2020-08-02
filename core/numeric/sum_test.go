package numeric

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Natural Numbers Sum", func(t *testing.T) {
		data := make([]int, 100)
		sum := float64(0)
		summer := NewSum()

		for i := 0; i <= len(data); i++ {
			summer.Add(i)
			sum += float64(i)
		}
		res, err := summer.Result()

		assert.Nil(t, err)
		assert.Equal(t, sum, res)
	})


	t.Run("Random Numbers Sum", func(t *testing.T) {
		data := make([]int, 1000)
		sum := float64(0)
		summer := NewSum()

		for i := 0; i <= len(data); i++ {
			v := (rand.Float64() - 0.5) * float64(i)
			summer.Add(v)
			sum += v
		}
		res, err := summer.Result()

		assert.Nil(t, err)
		assert.Equal(t, sum, res)
	})
}
