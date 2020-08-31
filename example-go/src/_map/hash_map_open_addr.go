package _map

import (
	"errors"
	"hash/crc32"
)

/**
	open addressing，也被称为封闭散列（closed hashing）实现HashMap
 */

const (
	DefaultSize       = 8    // 默认长度
	DefaultLoadFactor = 0.75 // 默认加载因子
)

type (
	// 节点数据
	Node struct {
		hash  int    // hash值
		key   string // 键
		value string // 值
		next  *Node  // 下一节点
	}

	// 字典对象
	HashMap struct {
		nodeArr    []Node  // 节点数据存储
		size       int     // 当前长度
		loadFactor float32 // 加载因子
	}
)

/***********************************************************************************************************
										node function
 ***********************************************************************************************************/

/**
	节点判断value是否一致，
 */
func (node *Node) IsEquals(value string) bool {
	// 如果长度不一致，则认定为值不匹配
	if len(node.value) != len(value) {
		return false
	}
	// 分解为byte数组，对同下标byte做异与操作，通过判断re是否有变化来计算是否一致，避免了定时攻击
	a, b, re := []byte(node.value), []byte(value), byte(0)
	for i := 0; i < len(value); i++ {
		re |= a[i] ^ b[i]
	}
	return re == byte(0)
}

/**
	获取hashcode
 */
func (node *Node) HashCode() int {
	return int(crc32.ChecksumIEEE([]byte(node.key))) ^ int(crc32.ChecksumIEEE([]byte(node.value)))
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

/***********************************************************************************************************
										public function
 ***********************************************************************************************************/

/**
	获取一个新的HashMap集合，长度为默认长度
 */
func NewHashMap() *HashMap {
	return &HashMap{
		nodeArr:    make([]Node, DefaultSize), // 构建一个默认长度的数组
		size:       DefaultSize,               // 设置默认长度
		loadFactor: DefaultLoadFactor,         // 设置默认加载因子
	}
}

/**
	获取一个新的Node节点对象
 */
func NewNode(key string, value string, hash int) *Node {
	return &Node{
		key:   key,   // 设置键
		value: value, // 设置值
		hash:  hash,  // 设置hash
		next:  nil,   // 将指针指向nil
	}
}

/**
	对key进行hash操作
 */
func hash(key string) int {
	h := int(crc32.ChecksumIEEE([]byte(key)))
	return h ^ (h >> 16)
}

/***********************************************************************************************************
										hashMap function
 ***********************************************************************************************************/
/**
	获取指定key对应的value
 */
func (hashMap *HashMap) Get(key string) (string, error) {
	// 如果当前数组为空或者map大小小于0，则返回空字符串
	if hashMap.nodeArr != nil && hashMap.size > 0 {
		kh := hash(key)                            // 获取查找key的hash值
		f := hashMap.nodeArr[kh&(hashMap.size-1) ] // 获取对应下标数据

		// 如果hash值匹配，直接返回对应数据，这里检测了node节点的hash值和node节点key的hash值
		if f.hash == kh &&
			(f.key == key || (key != "" && hash(f.key) == kh)) {
			return f.value, nil
		}

		// 如果链表无下一节点，则返回空
		if f.next == nil {
			return "", errors.New("is empty")
		}

		e, condition := f.next, true
		for ; condition; {
			// 判断hash是否一致，一致则返回对应value，hash检测逻辑同上
			if e.hash == kh &&
				(e.key == key || (key != "" && hash(e.key) == kh)) {
				return e.value, nil
			}
			e, condition = e.next, e.next != nil // 查找链表的下一值，这里如果无next，condition条件会终止for循环
		}
	}
	return "", errors.New("is empty")
}

func (hashMap *HashMap) Put(key string, value string) {

}

/**
	删除指定key
 */
func (hashMap *HashMap) Remove(key string) string {

	return ""
}

/**
	清空Map
 */
func (hashMap *HashMap) Clear() {
	if hashMap.nodeArr != nil && hashMap.size > 0 {
		hashMap.size = 0
		for n := range hashMap.nodeArr {
			// 遍历数组，将指针置为nil
			&n = nil
		}
	}
}

/**
	获取当前map长度
 */
func (hashMap *HashMap) Size() int {
	return hashMap.size
}

/**
	判断是否包含key
 */
func (hashMap *HashMap) ContainsKey(k string) bool {
	_, err := hashMap.Get(k)
	return err == nil
}

/**
	判断是否包含指定值
 */
func (hashMap *HashMap) ContainsValue(v string) bool {
	if hashMap.nodeArr != nil && hashMap.size > 0 {
		for i := 0; i < len(hashMap.nodeArr); i++ {
			for node := hashMap.nodeArr[i]; &node != nil; node = *node.next {
				if node.value == v ||
					(&node.value != nil && hash(node.value) == hash(v)) {
					return true
				}
			}
		}
	}
	return false
}

/**
	获取值数组
 */
func (hashMap *HashMap) Values() []string {
	re := make([]string, hashMap.size)
	for i, a := 0, 0; i < len(hashMap.nodeArr); i++ {
		for node := hashMap.nodeArr[i]; &node != nil; node = *node.next {
			re[a] = node.value
			a++
		}
	}
	return re
}
