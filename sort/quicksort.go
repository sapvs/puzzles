package sort

func QuickSort(nums []int) ([]int, SortStat) {
	n := len(nums)
	stat := SortStat{
		n:           n,
		comparisons: 0,
		swaps:       0,
	}
	quickSort(nums, 0, n-1, &stat)
	return nums, stat
}

func quickSort(nums []int, low, high int, stat *SortStat) {
	if low < high {
		part := partition(nums, low, high, stat)

		quickSort(nums, low, part-1, stat)
		quickSort(nums, part+1, high, stat)
	}

}
func partition(nums []int, low, high int, stat *SortStat) int {
	pivot := nums[high]
	i := low - 1

	for j := low; j < high; j++ {
		stat.comparisons++
		if nums[j] < pivot {
			i++
			swap(nums, i, j)
			stat.swaps++
		}
	}
	swap(nums, i+1, high)
	stat.swaps++
	return i + 1
}
