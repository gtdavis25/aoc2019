package main

type pathQueue []*path

func (q pathQueue) Len() int           { return len(q) }
func (q pathQueue) Less(i, j int) bool { return q[i].distance < q[j].distance }

func (q pathQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *pathQueue) Pop() interface{} {
	path := (*q)[len(*q)-1]
	path.index = -1
	*q = (*q)[:len(*q)-1]
	return path
}

func (q *pathQueue) Push(x interface{}) {
	path := x.(*path)
	path.index = len(*q)
	*q = append(*q, path)
}
