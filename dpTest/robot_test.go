package dpTest

import (
	"fmt"
	"testing"
)

func TestWay(t *testing.T) {
	if Way1(7, 4, 9, 5) == Way2(7, 4, 9, 5) {
		t.Logf("ok")
	}
	fmt.Println(Way1(5, 4, 4, 4))
	fmt.Println(Way3(5, 4, 4, 4))
}
