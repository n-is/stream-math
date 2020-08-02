package numeric

// Source: Online section of https://www.probabilitycourse.com/chapter5/5_3_1_covariance_correlation.php

type Covariance struct {
	mX, mY, mXY *Avg
}

func NewCovariance() *Covariance {

	cov := &Covariance{
		mX:  NewAvg(),
		mY:  NewAvg(),
		mXY: NewAvg(),
	}
	cov.Reset()

	return cov
}

func (m *Covariance) Add(x, y float64) {
	m.mX.Add(x)
	m.mY.Add(y)
	m.mXY.Add(x * y)
}

func (m *Covariance) Remove(x, y float64) {
	m.mX.Remove(x)
	m.mY.Remove(y)
	m.mXY.Remove(x * y)
}

func (m *Covariance) Result() (float64, error) {
	mx, _ := m.mX.Result()
	my, _ := m.mY.Result()
	mxy, _ := m.mXY.Result()

	return mxy - mx*my, nil
}

func (m *Covariance) Reset() {
	m.mX.Reset()
	m.mY.Reset()
	m.mXY.Reset()
}
