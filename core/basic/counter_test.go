package basic

import (
	"fmt"
	"math/rand"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestCounter (t *testing.T) {
	t.Run("Fixed Length Count", func(t *testing.T) {
		data := make([]int, 20)
		count := 20
		counter := New_Counter()

		for i := 0; i < len(data); i++ {
			counter.Count()
		}
		res := counter.Total_Count()

		assertEqual(t, count, res, "")
	})

	t.Run("Random Length Count", func(t *testing.T) {
		data := make([]int, (rand.Intn(50)))
		count := 0
		counter := New_Counter()

		for i := 0; i < len(data); i++ {
			counter.Count()
			count ++
		}
		res := counter.Total_Count()
		assertEqual(t, count, res, "Error!! Problem in counter")
	})

	t.Run("Conter Reset", func(t *testing.T) {
		data := []int{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		}
		reset := rand.Intn(10)
		count := 10-reset
		counter := New_Counter()

		for _,x := range data {
			counter.Count()
			if x == reset{
				counter.Reset()
			}
		}
		res := counter.Total_Count()

		assertEqual(t, count, res, "")
	})
}
