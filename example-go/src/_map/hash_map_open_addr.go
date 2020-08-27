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
		next  *Node
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
func NewHashMap(size int) *HashMap {
	return &HashMap{
		nodeArr: make([]Node, 0),
		size:    DefaultSize,
	}
}

/**
	获取一个新的Node节点对象
 */
func NewNode(key string, value string, hash int) *Node {
	return &Node{
		key:   key,
		value: value,
		hash:  hash,
		next:  nil,
	}
}

/**
	对key进行hash操作
 */
func hash(key string) int {
	h := int(crc32.ChecksumIEEE([]byte(key)))
	// fixme 搞懂为什么右移16位
	return h ^ (h >> 16)
}

/**
	获取指定key对应的value
 */
func (hashMap *HashMap) Get(key string) string {
	// 如果当前数组为空或者map大小小于0，则返回空字符串
	if hashMap.nodeArr != nil && hashMap.size > 0 {
		keyHash := hash(key)                  // 获取查找key的hash值
		index := keyHash & (hashMap.size - 1) // 计算hash值在数组中的下标
		first := hashMap.nodeArr[index]       // 获取对应下标数据

		// 如果hash值匹配，直接返回对应数据
		if first.hash == keyHash &&
			(first.key == key || (key != "" && hash(first.key) == keyHash)) {
			return first.value
		}

		// 向后查找链表的下一个值
		e := first.next
		// 如果无值，则返回空
		if e == nil {
			return ""
		}
		condition := true
		for ; condition; {
			// 判断hash是否一致，一致则返回对应value
			if e.hash == keyHash &&
				(e.key == key || (key != "" && hash(e.key) == keyHash)) {
				return e.value
			} else {
				// 继续向下获取链表
				e = e.next
			}
			// 判断节点是否为空，为空则终止for循环，返回默认值
			condition = e != nil
		}
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
