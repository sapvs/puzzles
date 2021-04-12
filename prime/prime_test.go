package prime

import (
	"fmt"
	"testing"
)

var funcs = []struct {
	name string
	f    func(int) bool
}{
	{name: "v1", f: PrimeV1},
	{name: "v2", f: PrimeV2},
	{name: "v3", f: PrimeV3},
}

func BenchmarkPrimePar(b *testing.B) {
	for _, ffn := range funcs {
		b.Run(fmt.Sprintf("%s-%d", ffn.name, b.N), func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					ffn.f(1000)
				}
			})
		})
	}
}

func BenchmarkPrimeSer(b *testing.B) {
	for _, ffn := range funcs {
		b.Run(fmt.Sprintf("%s-%d", ffn.name, b.N), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ffn.f(99999)
			}
		})
	}
}
