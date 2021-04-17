package prime

import (
	"fmt"
	"testing"
)

var testfuncs = []struct {
	name string
	f    func(int) bool
}{
	{name: "PrimeV1", f: IsPrimeV1},
	{name: "PrimeV2", f: IsPrimeV2},
	{name: "PrimeV3", f: IsPrimeV3},
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{9789, false},
	}

	for _, test := range testfuncs {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s-%d", test.name, tt.in), func(t *testing.T) {
				if got := test.f(tt.in); got != tt.want {
					t.Errorf("%s = %v, want %v", test.name, got, tt.want)
				}
			})
		}
	}
}

var b1 bool

func BenchmarkSer(b *testing.B) {
	for _, test := range testfuncs {
		b.Run(fmt.Sprintf("%s-%d", test.name, b.N), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b1 = test.f(999)
			}
		})
	}
}

func BenchmarkV1Ser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b1 = IsPrimeV1(999)
	}
}

func BenchmarkV2Ser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b1 = IsPrimeV2(999)
	}
}

func BenchmarkV3Ser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b1 = IsPrimeV3(999)
	}
}

func BenchmarkV1Par(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			IsPrimeV1(999)
		}
	})
}

func BenchmarkV2Par(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			IsPrimeV2(999)
		}
	})
}

func BenchmarkV3Par(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			IsPrimeV3(999)
		}
	})
}
