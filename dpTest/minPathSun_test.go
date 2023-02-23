package dpTest

import (
	"fmt"
	"testing"
)

func TestMinPathSum1(t *testing.T) {
	m := [][]int{
		{1, 3, 5, 9},
		{8, 1, 3, 4},
		{5, 0, 6, 1},
		{8, 8, 4, 0},
	}
	fmt.Println(MinPathSum1(m))
}
