package basic

type test interface {
	count()
	remove()
	Total_count() int
	Reset()
}

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

func (c *Counter) Total_Count() int{
	return c.counter
}

func (c *Counter) Reset() {
	c.counter = 0
}


