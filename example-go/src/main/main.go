package main

func main() {
	node := &Node{
		"string",
		"stirng",
		//nil,
	}
	result := node.isEquals("string")
	print(result)
}

/**
	放定址法（open addressing），也被称为封闭散列（closed hashing）实现HashMap
 */

const (
	DefaultSize = 8 // 默认长度
)

type (
	// 节点数据
	Node struct {
		key   string
		value string
		//next  Node
	}

	// 字典对象
	HashMap struct {
		nodeArr []Node
		size    int
	}
)


/**
	判断节点是否一致
 */
func (node *Node) isEquals(value string) bool {
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
	获取一个新的HashMap集合，长度为默认长度
 */
func New(size int) *HashMap {
	return &HashMap{
		nodeArr: make([]Node, DefaultSize),
		size:DefaultSize,
	}
}






