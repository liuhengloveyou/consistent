package consistent_test

import (
	"fmt"
	"testing"

	. "github.com/liuhengloveyou/consistent"
)

func TestNewConsistent(t *testing.T) {
	o := NewConsistent(nil)
	fmt.Println(o)

	o1 := NewConsistent([]string{"aaa","bbb","ccc"})
	fmt.Println(o1)
}
