package sort

func BubbleSort(nums []int) ([]int, SortStat) {
	n := len(nums)
	sortStat := SortStat{
		n:           n,
		comparisons: 0,
		swaps:       0,
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			sortStat.comparisons++
			if nums[j] > nums[j+1] {
				sortStat.swaps++
				swap(nums, j, j+1)
			}
		}
	}
	return nums, sortStat
}

func swap(nums []int, from, to int) {
	temp := nums[from]
	nums[from] = nums[to]
	nums[to] = temp
}
