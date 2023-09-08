package counter

type Counter map[string]int

func New() Counter {
	return make(map[string]int)
}

func (c Counter) Inc(str string) {
	b, ok := c[str]
	if !ok {
		b = 0
	}
	c[str] = b + 1
}
