package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := 7

	result := Sum(nums...)

	if expected != result {
		t.Error("Result should be 6")
	}
}

func TestSumFailRequire(t *testing.T) {
	num := []int{1, 2, 3}
	expected := 40
	result := Sum(num...)

	require.Equal(t, expected, result) // menghentikan semua dibawahnya
	fmt.Println("test fail sum terhenti")
}

func TestSumFailAssert(t *testing.T) {
	num := []int{1, 2, 3}
	expected := 40
	result := Sum(num...)

	assert.Equal(t, expected, result) // menjalankan dibawah nya
	fmt.Println("test fail sum terhenti")
}

func TestTableSum(t *testing.T) {
	testCases := []struct {
		result   int
		expected int
	}{
		{
			result:   Sum(1, 2, 3, 4),
			expected: 10,
		},
		{
			result:   Sum(1, 2, 3, 4),
			expected: 20,
		},
		{
			result:   Sum(1, 2, 3, 4),
			expected: 30,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, tC.expected, tC.result)
		})
	}
}
