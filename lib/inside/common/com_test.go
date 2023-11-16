package common

import (
	"fmt"
	"testing"
)

/*
   @Author: orbit-w
   @File: com_test
   @2023 11月 周三 17:37
*/

func TestPosNum_XY(t *testing.T) {
	p := &Pos{X: 500, Y: 892}
	n := p.Id()
	fmt.Println(n)
	fmt.Println(n.XY())

}
