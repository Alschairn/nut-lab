package _map

/**
	放定址法（open addressing），也被称为封闭散列（closed hashing）实现HashMap
 */

const (
	DefaultSize = 8 // 默认长度
)

type (
	// 节点数据
	Node struct {
		key   interface{}
		value interface{}
		next  Node
	}

	// 字典对象
	HashMap struct {
		nodeArr []Node
		size    int
	}
)


/**
	获取一个新的HashMap集合
 */
func New(size int) *HashMap {
	return &HashMap{
		nodeArr: make([]Node, size),
		size:0,
	}
}

