package main

type edgeQueue []edge

func (q edgeQueue) Len() int            { return len(q) }
func (q edgeQueue) Less(i, j int) bool  { return q[i].length < q[j].length }
func (q edgeQueue) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *edgeQueue) Push(x interface{}) { *q = append(*q, x.(edge)) }

func (q *edgeQueue) Pop() interface{} {
	item := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return item
}
