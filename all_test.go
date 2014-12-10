package consistent_test

import (
	"fmt"
	"testing"

	. "github.com/liuhengloveyou/consistent"
)

func TestNewConsistent(t *testing.T) {
	o := NewConsistent(nil, 0)
	fmt.Println(o)

	o1 := NewConsistent([]string{"aaa","bbb","ccc"}, 3)
	fmt.Println(o1)
}

func TestAdd(t *testing.T) {
	one := NewConsistent([]string{"aaa","bbb","ccc"}, 3)
	one.Add("zzzzz")
	t.Log(one)
}


func TestRemove(t *testing.T) {
	fmt.Println("\n\nremove>>>>>>>>>>>>>>>>>>>>>>>>>")
	one := NewConsistent([]string{"aaa","bbb","ccc"}, 3)
	one.Remove("zzzz")
	t.Log(one)
	
	fmt.Println("\n\nremove>>>>>>>>>>>>>>>>>>>>>>>>>1")
	one.Remove("aaa")
	t.Log(one)
}
