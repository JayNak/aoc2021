package util

type Stack []int

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s *Stack) Pop() int {
	l := len(*s)

	if l == 0 {
		panic("Empty Stack")
	}
	r := (*s)[l-1]
	*s = (*s)[:l-1]
	return r
}
