package sort

func InsertionSort(nums []int) ([]int, SortStat) {
	n := len(nums)
	sortStat := SortStat{
		n:           n,
		comparisons: 0,
		swaps:       0,
	}

	for i := 1; i < n; i++ {
		j := i - 1
		temp := nums[i]
		for j >= 0 && nums[j] > temp {
			sortStat.comparisons++
			sortStat.swaps++
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = temp
	}
	return nums, sortStat
}
