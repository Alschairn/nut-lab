package _map

/**
拉链法(separate chaining)，也称为开放散列法(open hashing)或者封闭定址法(closed hashing) 实现HashMap
*/

type (
	OpenHashingNode struct {
		hash  int    // hash值
		key   string // 键
		value string // 值
	}

	OpenHashMap struct {
		nodes      []OpenHashingNode
		size       int
		loadFactor float32 // 加载因子
		threshold  int     // 最大长度
	}
)
