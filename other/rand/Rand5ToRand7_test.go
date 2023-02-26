package rand

import (
	"fmt"
	"testing"
)

//  ok!
//  === RUN   TestRand1To5
//  5 3 5 5 1 0 1 2 4 0 --- PASS: TestRand1To5 (0.00s)
//  PASS

func TestRand1To5(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", Rand1To5())
	}
}

func TestRand01Fun(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", Rand01Fun())
	}
}

/*
=== RUN   TestRand1To7
1111736 1110610 1110942 1110508 1110616 1112188 1111220,1111068--- PASS: TestRand1To7 (0.85s)
PASS
*/
func TestRand1To7(t *testing.T) {
	var (
		a0 int
		a1 int
		a2 int
		a3 int
		a4 int
		a5 int
		a6 int
		a7 int
	)
	for i := 0; i < 8888888; i++ {
		x := Rand1To7()
		switch x {
		case 0:
			a0++
		case 1:
			a1++
		case 2:
			a2++
		case 3:
			a3++
		case 4:
			a4++
		case 5:
			a5++
		case 6:
			a6++
		case 7:
			a7++
		}

	}
	fmt.Printf("%d %d %d %d %d %d %d,%d", a0, a1, a2, a3, a4, a5, a6, a7)
}
