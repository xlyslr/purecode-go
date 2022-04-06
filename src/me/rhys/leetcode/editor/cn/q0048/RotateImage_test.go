package q0048

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	matrix1 := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	rotate(matrix1)
	fmt.Printf("%#v\n", matrix1)

	matrix2 := [][]int{
		{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16},
	}
	rotate(matrix2)
	fmt.Printf("%#v\n", matrix2)
}
