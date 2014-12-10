package consistent

import (
	"fmt"
	"hash/crc32"
	"sort"
)

// 环上的一个节点
type rnode struct {
	key uint32  // val 的哈希值
	val *string // 节点的值, 指向Consistent.members里的字段
}

type Consistent struct {
	ring    []rnode  // 哈希环
	members []string // 节点列表
	count   int      // 节点个数
	dup     int      // 节点重复次数
}

func NewConsistent(members []string, dup int) *Consistent {
	one := new(Consistent)
	one.members = members
	one.dup = dup
	one.count = len(members)
	one.ring = make([]rnode, one.count*one.dup)
	one.initRing()

	return one
}

func (this *Consistent) initRing() {
	p := 0

	for i := 0; i < this.count; i++ {
		for j := 0; j < this.dup; j++ {
			this.ring[p].key = this.hashKey(this.dupKey(this.members[i], j))
			this.ring[p].val = &this.members[i]
			p += 1
		}
	}

	sort.Sort(this)

	for i := 0; i < len(this.ring); i++ {
		fmt.Println(i, this.ring[i].key, *this.ring[i].val)
	}

	return
}

func (this *Consistent) Add(node string) {
	this.members = append(this.members, node)
	this.count += 1
	this.ring = make([]rnode, this.count*this.dup)
	this.initRing()
}

func (this *Consistent) Remove(node string) {
	p := sort.Search(len(this.members), func(i int) bool { return this.members[i] == node })
	if p < len(this.members) && this.members[p] == node {
		for ; p < len(this.members)-1; p++ {
			this.members[p] = this.members[p+1]
		}

		this.count -= 1
		this.ring = make([]rnode, this.count*this.dup)
		this.initRing()
	}
}

func (this *Consistent) Members() []string {
	return this.members
}

func (this *Consistent) hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (this *Consistent) dupKey(nodeVal string, idx int) string {
	return fmt.Sprintf("%s#%d", nodeVal, idx)
}

// 用来排序ring
func (this *Consistent) Len() int           { return len(this.ring) }
func (this *Consistent) Less(i, j int) bool { return this.ring[i].key < this.ring[j].key }
func (this *Consistent) Swap(i, j int)      { this.ring[i], this.ring[j] = this.ring[j], this.ring[i] }
