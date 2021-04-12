package other

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkRadn(b *testing.B) {

	fmt.Println(b.Name())
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}
