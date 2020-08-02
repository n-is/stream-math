package numeric

import (
	"github.com/n-is/stream-math/core"
	"github.com/n-is/stream-math/core/sort"
	"github.com/n-is/stream-math/utils/cast"
	"math"
)

type quantile struct {
	qth         float64
	fqCnt       *core.FreqCounter
	orderedVals sort.Sort
}

func newQuantileFunc(qth float64) *quantile {

	return &quantile{
		qth:   qth,
		fqCnt: core.NewFreqCounter(),
		orderedVals: sort.NewInsertion(func(old, new interface{}) bool {
			o, _ := cast.TryFloat(old)
			n, _ := cast.TryFloat(new)

			return o > n
		}),
	}
}

func (q *quantile) Add(v float64) {

	if vl := q.fqCnt.Add(v); vl != nil {
		q.orderedVals.Add(v)
	}

}

func (q *quantile) Remove(v float64) {

	if vl := q.fqCnt.Remove(v); vl != nil {
		q.orderedVals.Remove(v)
	}

}

func (q *quantile) Result() (float64, error) {
	res := q.quantile()
	return res, nil
}

func (q *quantile) quantile() float64 {
	qth := q.qth
	n := q.fqCnt.TotalFreq()

	i := qth * float64(n+1)
	j := math.Floor(i)

	values := q.fqCnt.Values()
	var jth, j1th interface{}

	cf := uint64(0)
	for e := q.orderedVals.First(); e != nil; e = e.Next() {
		entry := e.Value
		cf += values[entry]

		if (j + 1) <= float64(cf) {
			j1th = entry
			break
		}

		if j <= float64(cf) {
			jth = entry
			if i == j {
				jthf, _ := cast.TryFloat(jth)
				return jthf
			}
		}
	}

	if jth == nil {
		jth = j1th
	}

	jthf, _ := cast.TryFloat(jth)
	j1thf, _ := cast.TryFloat(j1th)

	return jthf + (j1thf-jthf)*(i-j)
}

func (q *quantile) Reset() {
	q.fqCnt.Reset()
	q.orderedVals.Reset()
}
