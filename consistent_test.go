package consistent_test

import (
	"fmt"
	"testing"

	. "github.com/liuhengloveyou/consistent"
)

func TestNewConsistent(t *testing.T) {
	fmt.Println("\n\nNewConsistent>>>>>>>>>>>>>>>>>>>>>>>>>")
	o := NewConsistent(nil, 0)
	fmt.Println(o)

	o1 := NewConsistent([]string{"zz","0000","aaa","bbb","ccc", "              "}, 3)
	fmt.Println(o1)
}

func TestInitRing(t *testing.T) {
	fmt.Println("\n\nInitRing>>>>>>>>>>>>>>>>>>>>>>>>>")

	o1 := NewConsistent([]string{"zz","0000","aaa","bbb","ccc", "      "}, 1)
	o1.InitRing()
	fmt.Println(o1)
}

func TestAdd(t *testing.T) {
	fmt.Println("\n\nAdd>>>>>>>>>>>>>>>>>>>>>>>>>")
	one := NewConsistent([]string{"aaa","bbb","ccc", "123"}, 3)
	fmt.Println("1>", one)
	one.Add("zzzzz")
	fmt.Println("2>", one)

	one.Add("")
	one.InitRing()
	fmt.Println("3>", one)
}


func TestRemove(t *testing.T) {
	fmt.Println("\n\nremove>>>>>>>>>>>>>>>>>>>>>>>>>")
	one := NewConsistent(nil, 3)
	one.Remove("zzz")
	fmt.Println("1>", one)

	one.InitRing()
	one.Remove("aaa")
	fmt.Println("2>", one)

	one.Add("aaa")
	one.Add("bbb")
	fmt.Println("3>", one)
	one.Remove("aaa")
	fmt.Println("3>", one)
}

func TestHash(t *testing.T) {
	fmt.Println("\n\nhash>>>>>>>>>>>>>>>>>>>>>>>>>")
	one := NewConsistent(nil, 3)
	fmt.Println(one.Hash("aaa"))

	one.Add("aaa")
	one.Add("bbb")
	one.Add("ccc")
	one.Add("1.1.1.1")
	one.InitRing()
	fmt.Println(one.Hash("aaa"))
	fmt.Println(one.Hash("aaa#1"))
	
}
