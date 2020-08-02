package numeric

import (
	"sync/atomic"
)

type Avg struct {
	mean float64
	num  uint64
}

func NewAvg() *Avg {
	return &Avg{
		mean: 0,
		num:  0,
	}
}

func (m *Avg) Add(v float64) {
	atomic.AddUint64(&m.num, 1)
	m.mean = m.mean + (v-m.mean)/float64(m.num)
}

func (m *Avg) Remove(v float64) {
	if m.num > 0 {
		m.num--
		if m.num == 0 {
			m.Reset()
		} else {
			m.mean = (m.mean*float64(m.num+1) - v) / float64(m.num)
		}
	}
}

func (m *Avg) Result() (float64, error) {
	return m.mean, nil
}

func (m *Avg) Reset() {
	m.mean = 0
	m.num = 0
}
