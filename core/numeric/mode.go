package numeric

import (
	"errors"
	"github.com/n-is/stream-math/core"
	"github.com/n-is/stream-math/core/sort"
	"github.com/n-is/stream-math/utils/cast"
)

type Mode struct {
	fqCnt       *core.FreqCounter
	first       bool
	orderedVals sort.Sort
}

func NewMode() *Mode {
	return &Mode{
		first: true,
		fqCnt: core.NewFreqCounter(),
		orderedVals: sort.NewInsertion(func(old, new interface{}) bool {
			o, _ := cast.TryFloat(old)
			n, _ := cast.TryFloat(new)

			return o > n
		}),
	}
}

func (md *Mode) Remove(v float64) {

	val, _ := cast.TryFloat(v)
	if v := md.fqCnt.Remove(val); v != nil {
		md.orderedVals.Remove(v)
	}

}

func (md *Mode) Add(v float64) {

	vl, _ := cast.TryFloat(v)
	if v := md.fqCnt.Add(vl); v != nil {
		md.orderedVals.Add(v)
	}

}

func (md *Mode) Result() (float64, error) {
	mode, err := md.calculate(md.fqCnt.Values())
	if err != nil {
		return 0, err
	}
	v, ok := mode.(float64)
	if !ok {
		return 0, errors.New("non-float mode")
	}
	return v, nil
}

func (md *Mode) Reset() {
	md.fqCnt.Reset()
}

func (md *Mode) calculate(m map[interface{}]uint64) (interface{}, error) {

	var mode interface{}
	mx := uint64(0)
	for e := md.orderedVals.First(); e != nil; e = e.Next() {
		k := e.Value
		v := m[k]
		if v > 0 {
			if v > mx {
				mx = v
				mode = k
			}
		}
	}

	return mode, nil
}
