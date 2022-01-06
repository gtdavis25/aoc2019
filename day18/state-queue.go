package main

type stateQueue []state

func (q stateQueue) Len() int            { return len(q) }
func (q stateQueue) Less(i, j int) bool  { return q[i].distance < q[j].distance }
func (q stateQueue) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *stateQueue) Push(x interface{}) { *q = append(*q, x.(state)) }

func (q *stateQueue) Pop() interface{} {
	item := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return item
}
