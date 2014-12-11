package consistent

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strings"
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
	one.dup = dup
	one.members = make([]string, 0)

	for i := 0; i < len(members); i++ {
		if strings.TrimSpace(members[i]) != "" {
			one.members = append(one.members, members[i])
		}
	}

	one.count = len(one.members)

	return one
}

// 初始化环. 更新节点信息要重新初始化环
func (this *Consistent) InitRing() error {
	p := 0

	if this.count < 1 {
		return fmt.Errorf("Init empty ring")
	}

	sort.Strings(this.members)                     // 排序节点
	this.ring = make([]rnode, this.count*this.dup) // 新的环

	for i := 0; i < this.count; i++ {
		for j := 0; j < this.dup; j++ {
			this.ring[p].key = this.hashKey(this.dupKey(this.members[i], j))
			this.ring[p].val = &this.members[i]
			p += 1 // GO的指针不能运算...
		}
	}

	sort.Sort(this) // 排序环

	/*
	for i := 0; i < len(this.ring); i++ {
		fmt.Println(this.ring[i].key, *this.ring[i].val)
	}*/

	return nil
}

// 添加一个节点. 不会重新初始化环
func (this *Consistent) Add(node string) error {
	if strings.TrimSpace(node) == "" {
		return fmt.Errorf("Add empty node.")
	}

	this.members = append(this.members, node)
	this.count += 1

	return nil
}

// 删除一个节点. 不会重新初始化环
func (this *Consistent) Remove(node string) {
	np, mlen := 0, len(this.members)

	p := sort.SearchStrings(this.members, node)
	if p < mlen && this.members[p] == node {
		tmp := this.members
		this.members = make([]string, mlen-1)
		for i := 0; i < mlen; i++ {
			if i != p {
				this.members[np] = tmp[i]
				np++
			}
		}

		this.count -= 1
	}

	return
}

// 重新设置重复次数
func (this *Consistent) SetDup(dup int) {
	this.dup = dup
}

// 取得所有节点列表
func (this *Consistent) Members() []string {
	return this.members
}

// 一致性哈希匹配(node <= key)
func (this *Consistent) Hash(key string) (string, error) {
	if strings.TrimSpace(key) == "" {
		return "", fmt.Errorf("Hash empty key.")
	}

	if this.count < 1 {
		return "", fmt.Errorf("empty ring")
	} else if this.count == 1 {
		return *this.ring[0].val, nil
	}

	keyHash := this.hashKey(key)
	// fmt.Println(">>>", keyHash)
	
	p := sort.Search(len(this.ring), func(i int) bool { return this.ring[i].key > keyHash })
	if p < len(this.ring) && this.ring[p].key > keyHash {
		if p == 0 {
			return *this.ring[len(this.ring)-1].val, nil
		} else {
			return *this.ring[p-1].val, nil
		}
	} else {
		return *this.ring[len(this.ring)-1].val, nil
	}

	return "", fmt.Errorf("service error")
}

// 算节点的哈希值
func (this *Consistent) hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

// 每个节点可以在环上出现多次
func (this *Consistent) dupKey(nodeVal string, idx int) string {
	return fmt.Sprintf("%s#%d", nodeVal, idx)
}

// 用来排序ring
func (this *Consistent) Len() int           { return len(this.ring) }
func (this *Consistent) Less(i, j int) bool { return this.ring[i].key < this.ring[j].key }
func (this *Consistent) Swap(i, j int)      { this.ring[i], this.ring[j] = this.ring[j], this.ring[i] }
