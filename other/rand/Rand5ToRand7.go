package rand

import "math/rand"

// this document is a series of function
//
// this function return int in [0,5]

func Rand1To5() int {
	return rand.Intn(6)
}

func Rand01Fun() (num int) {
	for {
		num = Rand1To5()
		if num == 1 || num == 2 {
			num = 0
			return
		}
		if num == 3 || num == 4 {
			num = 1
			return
		}
	}
	return
}

func Rand1To7() int {
	//  有0~5的随机函数，我们将其改造为0~1的随机函数
	var a int = 0
	for time := 0; time < 3; time++ {
		x := Rand01Fun()
		a = (a << 1) | x
	}
	return a
}
