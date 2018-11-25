package time

type Scalar struct {
	Id   int
	Time int
}

func (t *Scalar) IncrementAndGet() int {
	t.Time += 1
	return t.Time
}

func (t *Scalar) CompareIncrementAndGet(receivedTime int) int {
	maxTime := max(receivedTime, t.Time)
	time := maxTime + 1
	return time
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
