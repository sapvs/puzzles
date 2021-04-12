package sort

import "fmt"

type SortStat struct {
	n           int
	comparisons int
	swaps       int
}

func (s SortStat) String() string {
	return fmt.Sprintf("For %d elements %d comparisons and %d swaps ", s.n, s.comparisons, s.swaps)
}

type Sorter interface {
	sort([]int) SortStat
}
