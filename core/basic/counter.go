package basic

type Counter struct {
	counter int
}

func New_Counter() *Counter {
	return &Counter{
		counter: 0,
	}
}

func (c *Counter) Count() {
	c.counter += 1
}

func (c *Counter) Remove() {
	c.counter -= 1
}

func (c *Counter) Result() (interface{}, error){
	return c.counter, nil
}

func (c *Counter) Reset() {
	c.counter = 0
}


