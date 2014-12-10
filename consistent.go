package consistent

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// 环上的一个节点
type rnode struct {
	key uint32  // val 的哈希值
	val *string // 节点的值, 指向Consistent.members里的字段
}

type Consistent struct {
	count   int      // 节点个数
	ring    []rnode  // 哈希环
	members []string // 节点列表
}

func NewConsistent(members []string) *Consistent {
	return &Consistent{
		count:   0,
		ring:    make([]rnode),
		members: members }
}
