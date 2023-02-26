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
	var arr [8]int
	for i := 0; i < 888888888; i++ {
		x := Rand1To7()
		arr[x]++
	}
	fmt.Printf("%d %d %d %d %d %d %d,%d", arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7])
}
