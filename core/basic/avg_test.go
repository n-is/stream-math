package basic

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvg(t *testing.T) {

	t.Run("Natural Numbers Avg", func(t *testing.T) {
		// Generate data
		data := make([]float64, 50)
		for i := 0; i < len(data); i++ {
			data[i] = float64(i + 1)
		}

		// Find Calculated Avg
		sum := float64(0)
		for i := 0; i < len(data); i++ {
			sum += data[i]
		}
		calculatedAverage := sum / float64(len(data))

		// Find Stream Avg
		streamAverage := InitAvg()
		streamAverage.StreamAvg(data)

		// Compare the results
		assert.Equal(t, calculatedAverage, streamAverage.avg)
	})

	t.Run("Random Numbers Avg", func(t *testing.T) {
		// Generate data
		data := make([]float64, 50)
		for i := 0; i < len(data); i++ {
			data[i] = (rand.Float64() - 0.5) * float64(i)
		}

		// Find Calculated Avg
		sum := float64(0)
		for i := 0; i < len(data); i++ {
			sum += data[i]
		}
		calculatedAverage := sum / float64(len(data))

		// Find Stream Avg
		streamAverage := InitAvg()
		streamAverage.StreamAvg(data)

		// Compare the results
		assert.Equal(t, calculatedAverage, streamAverage.avg)
	})

}
