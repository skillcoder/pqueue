# pqueue

Priority queue built using the heap interface

## USING

```go
// Some items and their priorities.
items := map[int]int{
1: 10, 2: 20, 3: 5,
}

// Create a priority queue, put the items in it, and
// establish the priority queue (heap) invariants.
pq := make(Priority, len(items))
i := 0
pq.Push(&Item{value: 0, priority: 0})
for value, priority := range items {
pq[i] = &Item{
value:    value,
priority: priority,
index:    i,
}
i++
}
heap.Init(&pq)

// Take the items out; they arrive in decreasing priority order.

fmt.Println(heap.Pop(&pq).(*Item))
fmt.Println(heap.Pop(&pq).(*Item))
fmt.Println(heap.Pop(&pq).(*Item))
fmt.Println(heap.Pop(&pq).(*Item))
```