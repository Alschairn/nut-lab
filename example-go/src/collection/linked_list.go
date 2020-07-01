package collection

import "constant"

type LinkedListNode struct {
	data constant.Object // 节点数据
	next *LinkedListNode // 下一节点指针
}

type LinkedList struct {
	size int             // 长度
	head *LinkedListNode // 头节点
	tail *LinkedListNode // 尾节点
}

/**
	初始化空链表
 */
func (self *LinkedList) init() {
	self.size = 0
	self.head = nil
	self.tail = nil
}

/**
	向链表尾部追加元素
 */
func (self *LinkedList) append(node *LinkedListNode) bool {
	if node == nil { // 节点为空则返回插入失败
		return false
	}
	node.next = nil // 设置节点的next为nil

	if self.size == 0 { // 如果当前链表为空链表，将节点设置为头节点
		self.head = node
	} else { // 否则获取当前链表的尾节点，设置当前尾节点的next值指向追加节点
		oldTail := self.tail
		oldTail.next = node
	}
	self.tail = node // 设置链表的尾节点
	self.size++      // 长度+1
	return true
}

/**
	指定位置插入节点
 */
func (self *LinkedList) insert(i int, node *LinkedListNode) bool {
	if node == nil || i > self.size || self.size == 0 { // 插入节点为空或者长度大于当前链表长度或者当前链表为空，均返回插入失败
		return false
	}
	if i == 0 { // 如果插入位置为头节点
		node.next = self.head // 则将插入节点的next设置为头节点
		self.head = node      // 将头节点指向插入节点
	} else { // 其他位置
		preItem := self.head
		for j := 1; j < i; j++ { // 获取指定位置的前一个节点
			preItem = preItem.next
		}
		node.next = preItem.next // 将插入节点的next设置为原来这个位置的节点，相当于替换
		preItem.next = node      // 将指定位置前一节点的next设置为插入节点
	}
	self.size++
	return true
}

