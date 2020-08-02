package generic

import (
	"github.com/n-is/stream-math/core"
)

type Count struct {
	fqCnt *core.FreqCounter
}

func NewCount() *Count {
	return &Count{
		fqCnt: core.NewFreqCounter(),
	}
}

func (c *Count) Add(v interface{}) {
	c.fqCnt.Add(v)
}

func (c *Count) Remove(v interface{}) {
	c.fqCnt.Remove(v)
}

func (c *Count) Result() (interface{}, error) {
	cnt, err := c.calculate(c.fqCnt.Values())
	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func (c *Count) Reset() {
	c.fqCnt.Reset()
}

func (c *Count) calculate(m map[interface{}]uint64) (uint64, error) {
	cnt := uint64(0)

	for _, v := range m {
		cnt += v
	}

	return cnt, nil
}
