package big_o

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddItems(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{
			name:   "1",
			input:  1,
			output: 3,
		},
		{
			name:   "2",
			input:  2,
			output: 6,
		},
		{
			name:   "3",
			input:  3,
			output: 9,
		},
		{
			name:   "4",
			input:  4,
			output: 12,
		},
		{
			name:   "5",
			input:  5,
			output: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addItems(tt.input)
			assert.Equal(t, tt.output, result)
		})
	}
}

func BenchmarkAddItems(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := addItems(i)
		assert.NotNil(b, result)
		b.ReportAllocs()
	}
}
