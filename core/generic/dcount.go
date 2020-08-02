package generic

import "github.com/n-is/stream-math/core"

type DCount struct {
	fqCnt *core.FreqCounter
}

func NewDCount() *DCount {
	return &DCount{
		fqCnt: core.NewFreqCounter(),
	}
}

func (c *DCount) Add(v interface{}) {
	c.fqCnt.Add(v)
}

func (c *DCount) Remove(v interface{}) {
	c.fqCnt.Remove(v)
}

func (c *DCount) Result() (interface{}, error) {
	dcnt, err := c.calculate(c.fqCnt.Values())
	if err != nil {
		return 0, err
	}

	return dcnt, nil
}

func (c *DCount) Reset() {
	c.fqCnt.Reset()
}

func (c *DCount) calculate(m map[interface{}]uint64) (uint64, error) {
	dcnt := uint64(0)

	for _, v := range m {
		if v > 0 {
			dcnt++
		}
	}

	return dcnt, nil
}
