package main

/*
	This algorithm solve the following problem:
    Find the independent set of intervals such that its cardinality is maximum.
	Input: N intervals such that I1=(a1,b1[ , ... , In=(an,bn[.
	Output: a set of maximum cardinality such that every interval not overlap
	        with any other.
*/

// This solution use a greedy approach

import (
	"container/list"
	"fmt"
	"sort"
)

// Interval structure to save the interval start and ends
type Interval struct {
	Start int
	End   int
}

// pretty print the interval
func (this Interval) String() string {
	return fmt.Sprintf("(%d, %d)", this.Start, this.End)
}

type ByEndTime []Interval

func (a ByEndTime) Len() int           { return len(a) }
func (a ByEndTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByEndTime) Less(i, j int) bool { return a[i].End < a[j].End }

// This function take a slice of Pair and compute the independent set of
// Pair, returned as *list.List.
// The greedy choice is made by ordering the given slice of Pair by ending time.
// A copy of the argument is done because sort.Sort() modify the input.
// Complexity: O( n + n*log(n) + n ) = O(n*log(n))
func IndependentSet(ranges []Interval) *list.List {
	intervals := make([]Interval, len(ranges))
	copy(intervals, ranges)
	sort.Sort(ByEndTime(intervals))

	independentSet := list.New()
	independentSet.PushBack(intervals[0])
	last := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start >= intervals[last].End {
			independentSet.PushBack(intervals[i])
			last = i
		}
	}

	return independentSet
}

func main() {
	// Intervals slice.
	intervals := []Interval{
		{2, 5},
		{1, 3},
		{2, 7},
		{4, 6},
		{4, 8},
		{2, 10},
	}

	fmt.Println("Set given:")
	fmt.Println(intervals)
	fmt.Println("Independent set of maximum cardinality:")
	independentSet := IndependentSet(intervals)
	for e := independentSet.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
