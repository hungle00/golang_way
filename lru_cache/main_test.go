package main

import "testing"

func TestNewLRUCache(t *testing.T) {
	// Test case 1: Create a new LRU cache with size 3
	cache := NewLRUCache(3)

	// Check if the cache initialized correctly
	if cache.capacity != 3 || cache.head == nil || cache.tail == nil || cache.head.next != cache.tail || cache.tail.prev != cache.head {
		t.Error("NewLRUCache initialization failed.")
	}

	// Test case 2: Create a new LRU cache with size 0
	emptyCache := NewLRUCache(0)

	// Check if the cache is initialized correctly for size 0
	if emptyCache.capacity != 0 || emptyCache.head == nil || emptyCache.tail == nil || emptyCache.head.next != emptyCache.tail || emptyCache.tail.prev != emptyCache.head {
		t.Error("NewLRUCache initialization failed for size 0.")
	}
}

func TestCacheMoveToHead(t *testing.T) {
	// Test case 1: Move a node to the head in a non-empty cache
	cache := NewLRUCache(3)

	// Add nodes to the cache
	node1 := &Node{value: 101, key: "a"}
	node2 := &Node{value: 102, key: "b"}
	node3 := &Node{value: 103, key: "c"}
	cache.addNode(node1)
	cache.addNode(node2)
	cache.addNode(node3)

	// Move node2 to the head
	cache.moveToHead(node2)

	// Check if node2 is correctly moved to the head
	if cache.head.next != node2 || node2.prev != cache.head || node2.next != node3 || node3.prev != node2 {
		t.Error("MoveToHead failed for a non-empty cache.")
	}

	// Test case 2: Move the last node to the head
	cache.moveToHead(node1)

	// Check if node1 has been relocated to the head and the tail is updated
	if cache.head.next != node1 || node1.prev != cache.head || node1.next != node2 || cache.tail.prev != node3 {
		t.Error("MoveToHead failed for the last node in the cache.")
	}

	// Test case 3: Move a node in a cache with only one node
	cache.moveToHead(node3)

	// Check if node3 is correctly moved to the head, and the cache remains unchanged
	if cache.head.next != node3 || node3.prev != cache.head || node3.next != node1 || cache.tail.prev != node2 {
		t.Error("MoveToHead failed for a cache with only one node.")
	}

	// Test case 4: Move a node that is not in the cache
	notInCache := &Node{value: 104, key: "d"}
	cache.moveToHead(notInCache)

	// Check if the cache remains unchanged when trying to move a node not in the cache
	if cache.head.next != node3 || node3.prev != cache.head || node3.next != node1 || cache.tail.prev != node2 {
		t.Error("MoveToHead failed for a node not in the cache.")
	}
}


func TestCachePopTail(t *testing.T) {
	// Test case 1: Pop the tail node in a non-empty cache
	cache := NewLRUCache(3)

	// Add nodes to the cache
	node1 := &Node{value: 101, key: "a"}
	node2 := &Node{value: 102, key: "b"}
	node3 := &Node{value: 103, key: "c"}
	cache.addNode(node1)
	cache.addNode(node2)
	cache.addNode(node3)

	// Pop the tail node
	poppedNode := cache.popTail()

	// Check if the tail node is correctly popped and returned
	if poppedNode != node1 || cache.tail.prev != node2 || node2.next != cache.tail {
		t.Error("PopTail failed for a non-empty cache.")
	}

	// Test case 2: Pop the tail node in a cache with only one node
	cache = NewLRUCache(1)
	soleNode := &Node{value: 201, key: "d"}
	cache.addNode(soleNode)

	// Pop the tail node
	poppedNode = cache.popTail()

	// Check if the sole node is correctly popped and returned, and the cache becomes empty
	if poppedNode != soleNode || cache.head.next != cache.tail || cache.tail.prev != cache.head {
		t.Error("PopTail failed for a cache with only one node.")
	}

	// Test case 3: Pop the tail node from an empty cache
	emptyCache := NewLRUCache(0)

	// Pop the tail node
	poppedNode = emptyCache.popTail()

	// Check if nil is returned when popping from an empty cache
	if poppedNode != nil {
		t.Error("PopTail failed for an empty cache.")
	}
}

func TestCacheGet(t *testing.T) {
	// Test case 1: Get a node with an existing key in the cache
	cache := NewLRUCache(3)

	// Add nodes to the cache
	node1 := &Node{value: 101, key: "a"}
	node2 := &Node{value: 102, key: "b"}
	node3 := &Node{value: 103, key: "c"}
	cache.addNode(node1)
	cache.items[node1.key] = node1
	cache.addNode(node2)
	cache.items[node2.key] = node2
	cache.addNode(node3)
	cache.items[node3.key] = node3

	// Get node with key "b"
	resultNode := cache.Get("b")

	// Check if the correct node is returned and moved to the head
	if resultNode != node2 || cache.head.next != node2 || node2.prev != cache.head || node2.next != node3 {
		t.Error("Get failed for an existing key in the cache.")
	}

	// Test case 2: Get a node with a non-existing key in the cache
	nonExistingKey := "non-existing"
	resultNode = cache.Get(nonExistingKey)

	// Check if nil is returned for a non-existing key, and the cache remains unchanged
	if resultNode != nil || cache.head.next != node2 || node2.prev != cache.head || node2.next != node3 {
		t.Error("Get failed for a non-existing key in the cache.")
	}
}

func TestCachePut(t *testing.T) {
	// Test case 1: Put a new node into the cache
	cache := NewLRUCache(3)

	// Put a new node with key "1" and value 101
	cache.Put("1", 101)

	// Check if the node is correctly added to the cache and hashMap
	if len(cache.items) != 1 || cache.head.next.key != "1" || cache.head.next.value != 101 || cache.tail.prev.key != "1" || cache.tail.prev.value != 101 {
		t.Error("Put failed for adding a new node.")
	}

	// Test case 2: Put a node with an existing key into the cache
	cache.Put("1", 102)

	// Check if the existing node is correctly updated and moved to the head
	if len(cache.items) != 1 || cache.head.next.key != "1" || cache.head.next.value != 102 || cache.tail.prev.key != "1" || cache.tail.prev.value != 102 {
		t.Error("Put failed for updating an existing node.")
	}

	// Test case 3: Put nodes into the cache until reaching capacity, causing eviction
	cache.Put("2", 202)
	cache.Put("3", 303)
	cache.Put("4", 404)

	// Check if the eviction occurred and the cache is at its capacity
	if len(cache.items) != 3 || cache.head.next.key != "4" || cache.head.next.value != 404 || cache.tail.prev.key != "2" || cache.tail.prev.value != 202 {
		t.Error("Put failed for reaching cache capacity and eviction.")
	}
}

