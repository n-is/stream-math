package basic

import (
	"github.com/n-is/stream-math/utils/cast"
)

// Avg type
type Avg struct {
	avg float64
}

// InitAvg function
func InitAvg() *Avg {
	return &Avg{
		avg: 0,
	}
}

// GetAvg function
func GetAvg(preAvg float64, v interface{}, n float64) float64 {
	newAvg := preAvg

	if val, ok := cast.TryFloat(v); ok {
		newAvg = (preAvg*n + val) / (n + 1)
	}

	return newAvg
}

// StreamAvg method
func (a *Avg) StreamAvg(arr []float64) {
	a.avg = 0

	for i := 0; i < len(arr); i++ {
		a.avg = GetAvg(a.avg, arr[i], float64(i))
	}

}

// func main() {
// 	fmt.Println("Beginning ... ")

// 	var numberStream = []float64{1000, 2, 3, 17, 50}
// 	a := InitAvg()

// 	a.StreamAvg(numberStream)
// 	fmt.Println("Stream Average is :", a.avg)

// 	fmt.Println("End ... ")
// }
