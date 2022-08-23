package errorx

import (
	"fmt"
	"testing"
)

func TestError(*testing.T) {
	e := New("自定义错误")
	fmt.Println(e)
}
