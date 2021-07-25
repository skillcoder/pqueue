// Package pqueue - a priority queue built using the heap interface.
package pqueue

import (
	"container/heap"
)

var _ heap.Interface = (*Priority)(nil)

// An Item is something we manage in a priority queue.
type Item struct {
	value    uint32 // The value of the item; arbitrary.
	priority int64  // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewItem(value uint32, priority int64) *Item {
	return &Item{
		value:    value,
		priority: priority,
		index:    0,
	}
}

func (item *Item) Value() uint32 {
	return item.value
}

// A Priority implements heap.Interface and holds Items.
type Priority []*Item

func (pq Priority) Len() int { return len(pq) }

func (pq Priority) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq Priority) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *Priority) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *Priority) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *Priority) update(item *Item, value uint32, priority int64) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
