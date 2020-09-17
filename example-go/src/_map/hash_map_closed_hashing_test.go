package _map

import "testing"

func TestNewNode(t *testing.T) {
	node := NewNode("key", "value", 1)
	isEquals := node.IsEquals("value")
	println(isEquals)
}

func TestNewHashMap(t *testing.T) {
	hashMap := NewHashMap()
	//println("size:", hashMap.size)
	//
	//hashMap.Put("key", "value")
	//v, _ := hashMap.Get("key")
	//println("value:", v)
	//println("arr len:", len(hashMap.nodeArr))
	//
	println("init 100 element")
	for i:=0; i<100; i++ {
		hashMap.Put(string(rune(i)), string(rune(i)))
	}
	println("size:", hashMap.size)
	println("arr len:", len(hashMap.nodeArr))


	for i:=100; i>0; i-- {
		v, _ := hashMap.Get(string(rune(i)))
		println(v)
	}
}