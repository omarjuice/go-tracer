package main

import (
	"container/heap"
	"math"
	"strconv"
)

//EPSILON is equivalence tolerance for float value comparison
const EPSILON = 0.00001
const π = math.Pi

//floatEqual determines if two floats are equal within a tolerance
func floatEqual(a, b float64) bool {
	return abs(a-b) < EPSILON
}
func floatToUint8String(f float64) string {
	if f < 0.0 {
		return "0"
	}
	f *= 256.0
	if f > 255.0 {
		return "255"
	}
	return strconv.Itoa(int(f))
}

//Abs returns absolute value
func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n

}

//Sum adds a sequence together
func sum(nums []float64) float64 {
	result := 0.0

	for _, v := range nums {
		result += v
	}
	return result
}

//A[i] * B[i] + A[i + 1] * B[i + 1] ...
func zipSum(A, B []float64) float64 {

	ln := int(min(float64(len(A)), float64(len(B))))
	total := 0.0
	for i := 0; i < ln; i++ {
		total += A[i] * B[i]
	}
	return total
}

//FloatToString converts a float to a String
func floatToString(n float64, cut int) string {
	// to convert a float number to a string
	s := strconv.FormatFloat(n, 'f', 6, 64)
	if cut > len(s) {
		return s[:]
	}
	return s[:cut]
}

func min(a, b float64) float64 {
	if b < a {
		return b
	}
	return a
}
func max(a, b float64) float64 {
	if b > a {
		return b
	}
	return a
}

//PriorityQueue of intersections
type PriorityQueue []*Intersection

//Len gets length
func (pq PriorityQueue) Len() int { return len(pq) }

//Less is comparator
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	a, b := pq[i].t, pq[j].t

	if a < 0 || b < 0 {
		return a > b
	}
	return a < b

}

//Swap swaps two items
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

//Push adds an item to the pq
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Intersection)
	item.index = n
	*pq = append(*pq, item)
}

//Pop removes the top item
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Intersection in the queue.
func (pq *PriorityQueue) update(intersection *Intersection, t float64) {
	intersection.t = t
	heap.Fix(pq, intersection.index)
}

func (pq *PriorityQueue) push(intersection *Intersection) int {
	heap.Push(pq, intersection)
	return pq.Len()
}

func (pq *PriorityQueue) pop() *Intersection {
	return heap.Pop(pq).(*Intersection)
}

//Empty tells if the PQ is empty
func (pq *PriorityQueue) Empty() bool {
	return pq.Len() == 0
}

//Top returns the top element
func (pq PriorityQueue) Top() *Intersection {
	return pq[0]
}

//Init to initialize
func (pq *PriorityQueue) Init() {
	heap.Init(pq)
}
