package _map

import "hash/crc32"

/**
	放定址法（open addressing），也被称为封闭散列（closed hashing）实现HashMap
	fixme 先搞string map，对象等到清楚泛型机制再搞
 */

const (
	DefaultSize = 8 // 默认长度
)

type (
	// 节点数据
	Node struct {
		hash  int
		key   string
		value string
		next  Node
	}

	// 字典对象
	HashMap struct {
		nodeArr []Node
		size    int
	}
)

/**
	节点判断value是否一致，
 */
func (node *Node) IsEquals(value string) bool {
	if len(node.value) != len(value) {
		return false
	}
	_1 := []byte(node.value)
	_2 := []byte(value)
	length := len(node.value)

	result := byte(0)
	for i := 0; i < length; i++ {
		result |= _1[i] ^ _2[i]
	}
	return result == byte(0)
}

/**
	获取hashcode
 */
func (node *Node) HashCode() int {
	key := int(crc32.ChecksumIEEE([]byte(node.key)))
	value := int(crc32.ChecksumIEEE([]byte(node.value)))
	if -key >= 0 {
		key = -key
	}
	if -value >= 0 {
		value = -value
	}
	return key ^ value
}

func (node *Node) SetValue(value string) string {
	oldValue := node.value
	node.value = value
	return oldValue
}

func (node *Node) GetKey() string {
	return node.key
}

func (node *Node) GetValue() string {
	return node.value
}

func (node *Node) ToString() string {
	return node.key + "=" + node.value
}

/**
	获取一个新的HashMap集合，长度为默认长度
 */
func New(size int) *HashMap {
	return &HashMap{
		nodeArr: make([]Node, 0),
		size:    DefaultSize,
	}
}

/**
	对key进行hash操作
 */
func hash(key string) int {
	h := int(crc32.ChecksumIEEE([]byte(key)))
	// fixme 搞懂为什么这么操作
	return h ^ (h >> 16)
}

func (hashMap *HashMap) Get(key string) string {
	if hashMap.nodeArr != nil && hashMap.size > 0 {
		hash := hash(key)
		index := hash & (hashMap.size - 1)
		first := hashMap.nodeArr[index]
		if first.hash == hash {
			return first.value
		}
		if first.next != nil

	}


	return ""
}

func (hashMap *HashMap) Put(key string, value string) {

}

func (hashMap *HashMap) PutAll(tMap HashMap) {

}

func (hashMap *HashMap) Remove(key string) string {
	return ""
}

func (hashMap *HashMap) Clear() {

}

func (hashMap *HashMap) Size() int {
	return hashMap.size
}

func (hashMap *HashMap) ContainsKey(key string) bool {
	return false
}

func (hashMap *HashMap) ContainsValue(value string) bool {
	return false
}

func (hashMap *HashMap) Values() []string {
	return nil
}
