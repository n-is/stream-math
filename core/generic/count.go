package generic

import (
	"sync/atomic"
)

type Count struct {
	num uint64
}

func NewCount() *Count {
	return &Count{
		num: 0,
	}
}

func (m *Count) Add(v interface{}) {
	atomic.AddUint64(&m.num, 1)
}

func (m *Count) Remove(v interface{}) {
	if m.num > 0 {
		m.num--
	}
}

func (m *Count) Result() (interface{}, error) {
	return m.num, nil
}

func (m *Count) Reset() {
	m.num = 0
}
