package myGin

import (
	"fmt"
	"testing"
)

func TestShouldOnlyGet(t *testing.T) {
	fn := "update"
	res := ShouldOnlyGet(fn)
	fmt.Println(res)
}
