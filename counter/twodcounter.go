package counter

type TwoDCounter map[string]map[string]int

func NewTwoDCounter() TwoDCounter {
	return make(map[string]map[string]int)
}

func (c TwoDCounter) Inc(str1, str2 string) {
	m, ok := c[str1]
	if !ok {
		m = make(map[string]int)
		c[str1] = m
	}

	i, ok := m[str2]
	if !ok {
		i = 0
	}

	m[str2] = i + 1
}
