package main

import (
	"fmt"
	"github.com/n-is/stream-math/core/numeric"
)

func main() {
	sum := numeric.NewSum()

	data := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for _, d := range data {
		sum.Add(d)
	}

	if res, err := sum.Result(); err == nil {
		fmt.Println("Sum of provided numbers:", res)
	}
}
