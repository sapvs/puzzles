package sort

func SelectionSort(nums []int) ([]int, SortStat) {
	n := len(nums)
	stat := SortStat{
		n:           n,
		swaps:       0,
		comparisons: 0,
	}

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if nums[min] > nums[j] {
				stat.comparisons++
				min = j
			}
		}
		if min != i {
			swap(nums, min, i)
			stat.swaps++
		}
	}
	return nums, stat
}
