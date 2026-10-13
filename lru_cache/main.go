package main

import "fmt"

type Node struct {
	value int
	key   string
	next  *Node
	prev  *Node
}

type LRUCache struct {
	capacity int
	items    map[string]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(size int) *LRUCache {
	headNode := &Node{prev: nil}
	tailNode := &Node{prev: headNode, next: nil}
	headNode.next = tailNode

	return &LRUCache{
		capacity: size,
		items   : make(map[string]*Node),
		head    : headNode,
		tail    : tailNode,
	}
}

func (cache *LRUCache) addNode(n *Node) {
	nextNode := cache.head.next
	prevNode := cache.head
	nextNode.prev = n
	n.next = nextNode
	n.prev = prevNode
	prevNode.next = n
}

func (cache *LRUCache) removeNode(n *Node) bool {
	nextNode := n.next
	prevNode := n.prev
    if prevNode == nil && nextNode == nil {
        return false
    }
	prevNode.next = nextNode
	nextNode.prev = prevNode

	return true
}

func (cache *LRUCache) moveToHead(n *Node) {
	if ok := cache.removeNode(n); ok {
		cache.addNode(n)
	}
}


func (cache *LRUCache) Get(key string) *Node {
	if node, exists := cache.items[key]; exists {
		cache.moveToHead(node)
		return node
	}
	return nil
}

func (cache *LRUCache) popTail() *Node {
	res := cache.tail.prev
    if res.prev == nil {
        return nil
    }
    cache.removeNode(res)
    return res
}

func (cache *LRUCache) Put(key string, value int) {
	if node, ok := cache.items[key]; ok {
		cache.removeNode(node)
	}
	newNode := &Node{
		value: value,
		key: key,
	}
	cache.addNode(newNode)
    cache.items[key] = newNode
    if len(cache.items) > cache.capacity {
        nodeToRemove := cache.popTail()
        delete(cache.items, nodeToRemove.key)
    } 
}

func main() {
	cache := NewLRUCache(3)
    cache.Put("a", 10)
    cache.Put("b", 20)
	cache.Put("c", 30)
	val := cache.Get("a") 
	cache.Put("d", 40)
	fmt.Println("LRU Cache created with capacity:", cache.capacity)
	fmt.Println(val.value) // Output: 10
}

