package big_o

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr    []int
		value  int
		result int
	}{
		{[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, 12, 5},
		{[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, 2, 0},
		{[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, 20, 9},
		{[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, 5, -1},
		{[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, 21, -1},
	}

	for _, test := range tests {
		result := binarySearch(test.arr, test.value)
		assert.Equal(t, test.result, result)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	value := 12

	for i := 0; i < b.N; i++ {
		binarySearch(arr, value)
		b.ReportAllocs()
	}
}
