package basic

import "github.com/n-is/stream-math/utils/cast"

type Sum struct {
	sum float64
}

func NewSum() *Sum {
	return &Sum{
		sum: 0,
	}
}

func (s *Sum) Add(v interface{}) {
	if val, ok := cast.TryFloat(v); ok {
		s.sum += val
	}
}

func (s *Sum) Remove(v interface{}) {
	if val, ok := cast.TryFloat(v); ok {
		s.sum -= val
	}
}

func (s *Sum) Result() (interface{}, error) {
	return s.sum, nil
}

func (s *Sum) Reset() {
	s.sum = 0
}
