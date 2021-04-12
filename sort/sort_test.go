package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

const N int = 1000

var test_table = []struct {
	name string
	f    func([]int) ([]int, SortStat)
}{
	{name: "Quick", f: QuickSort},
	{name: "Selection", f: SelectionSort},
	{name: "Insertion", f: InsertionSort},
	{name: "Bubble", f: BubbleSort},
}

func createData() []int {
	nums := make([]int, 0)
	for i := N; i > 0; i-- {
		nums = append(nums, rand.Intn(100))
	}
	return nums
}

func print(arr []int) {
	for _, val := range arr {
		fmt.Printf("%d \t", val)
	}
	fmt.Println()
}

func verify(nums []int, t *testing.T) {
	n := len(nums)
	if n != N {
		t.Fail()
	}
	for i := 0; i < N-1; i++ {
		if nums[i] > nums[i+1] {
			t.Fail()
		}
	}

}

func TestSorts(t *testing.T) {
	for _, test := range test_table {
		ret, stat := test.f(createData())
		verify(ret, t)
		t.Logf("%s Stats : %v", test.name, stat)
	}
}

func BenchmarkParallel(b *testing.B) {
	for _, test := range test_table {
		b.Run(test.name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					test.f(createData())
				}
			})
		})
	}
}

func BenchmarkSerial(b *testing.B) {
	for _, test := range test_table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i <= b.N; i++ {
				test.f(createData())
			}
		})
	}
}
