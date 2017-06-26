// Code listing 53: https://play.golang.org/p/8_S4aUUD89

package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

// following are for container/heap
type OrderedNums []int

func (h OrderedNums) Len() int {
	return len(h)
}
func (h OrderedNums) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h OrderedNums) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *OrderedNums) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *OrderedNums) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}
// end declarations for container/heap

func main() {
	// *** container/list ***
	l := list.New()
	e0 := l.PushBack(42)
	e1 := l.PushFront(11)
	e2 := l.PushBack(19)
	l.InsertBefore(7, e0)
	l.InsertAfter(254, e1)
	l.InsertAfter(4987, e2)

	fmt.Println("*** LIST ***")
	fmt.Println("-- Step 1:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value.(int))
	}
	fmt.Printf("\n")

	l.MoveToFront(e2)
	l.MoveToBack(e1)
	l.Remove(e0)

	fmt.Println("-- Step 2:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value.(int))
	}
	fmt.Printf("\n")

	// *** container/ring ***
	// create the ring and populate it
	blake := []string{"the", "invisible", "worm"}

	r := ring.New(len(blake))
	for i := 0; i < r.Len(); i++ {
		r.Value = blake[i]
		r = r.Next()
	}

	// move (2 % r.Len())=1 elements forward in the ring
	r = r.Move(2)

	fmt.Printf("\n*** RING ***\n")
	// print all the ring values with ring.Do()
	r.Do(func(x interface{}) {
		fmt.Printf("%s\n", x.(string))
	})

	// *** container/heap
	h := &OrderedNums{34, 24, 65, 77, 88, 23, 46, 93}

	heap.Init(h)

	fmt.Printf("\n*** HEAP ***\n")
	fmt.Printf("min: %d\n", (*h)[0])
	fmt.Printf("heap:\n")
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Printf("\n")
}