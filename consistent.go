package consistent

import (
	"error"
	"fmt"
)

// 环上的一个节点
type rnode struct {
	key uint32  // val 的哈希值
	val *string // 节点的值, 指向Consistent.members里的字段
}

type Consistent struct {
	count   int      // 节点个数
	dup     int      // 节点重复次数
	ring    []rnode  // 哈希环
	members []string // 节点列表
}

func NewConsistent(members []string, dup int) *Consistent {
	return &Consistent{
		count:   0,
		dup:     dup,
		ring:    make([]rnode, len(members)),
		members: members}
}

func (this *Consistent) InitRing() error {

}

func (this *Consistent) dupKey(nodeVal string, idx int) string {
	return fmt.Sprintf("%s#%d", nodeVal, idx)
}
