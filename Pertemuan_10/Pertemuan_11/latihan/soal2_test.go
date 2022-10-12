package latihan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPetarung(t *testing.T) {
	testCases := []struct {
		petarung []int
		kekuatan []int
		self     int
		result   int
		expected int
	}{
		{
			petarung: []int{5, 3, 9, 8},
			kekuatan: []int{2, 2, 3, 1},
			self:     3,
			expected: 7,
		},
		{
			petarung: []int{2, 6, 3, 9},
			kekuatan: []int{2, 2, 3, 5},
			self:     2,
			expected: 14,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			assert.Equal(t, tC.expected, Pertarungan(tC.petarung, tC.kekuatan, tC.self))
		})
	}
}
