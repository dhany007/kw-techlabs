package latihan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPria(t *testing.T) {
	testCases := []struct {
		result   bool
		expected bool
	}{
		{
			result:   IsAllowWork("Pria", 16),
			expected: false,
		},
		{
			result:   IsAllowWork("Pria", 17),
			expected: false,
		},
		{
			result:   IsAllowWork("Pria", 50),
			expected: true,
		},
		{
			result:   IsAllowWork("Pria", 60),
			expected: false,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			assert.Equal(t, tC.expected, tC.result)
		})
	}
}

func TestWanita(t *testing.T) {
	testCases := []struct {
		result   bool
		expected bool
	}{
		{
			result:   IsAllowWork("Wanita", 16),
			expected: false,
		},
		{
			result:   IsAllowWork("Wanita", 21),
			expected: true,
		},
		{
			result:   IsAllowWork("Wanita", 50),
			expected: true,
		},
		{
			result:   IsAllowWork("Wanita", 60),
			expected: false,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintf("test - %d", i), func(t *testing.T) {
			assert.Equal(t, tC.expected, tC.result)
		})
	}
}
