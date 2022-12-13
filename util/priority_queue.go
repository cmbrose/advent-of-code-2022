package util

import "container/heap"

type PriorityQueue[T any] struct {
	arr  []T
	less func(a, b T) bool
}

func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		less: less,
	}
}

var _ heap.Interface = &PriorityQueue[struct{}]{}

func (pq *PriorityQueue[T]) Push(x any) {
	pq.arr = append(pq.arr, x.(T))
}

func (pq *PriorityQueue[T]) Pop() any {
	lastIdx := len(pq.arr) - 1

	item := pq.arr[lastIdx]
	pq.arr = pq.arr[:lastIdx]

	return item
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.arr)
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	a, b := pq.arr[i], pq.arr[j]

	return pq.less(a, b)
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}
