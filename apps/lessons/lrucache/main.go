package main

type Node struct {
	key   string
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {

	// If we add more than this, we evict the least recently used node, i.e. the tail.
	capacity int

	head *Node
	tail *Node

	// map for O(1) lookup. This has to be manually maintained as we add and remove nodes.
	cache map[string]*Node
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*Node),
	}
}

func (c *LRUCache) Put(key string, value int) {

	if len(c.cache) == 0 {
		newNode := &Node{
			key:   key,
			value: value,
		}
		c.cache[key] = newNode
		c.head = newNode
		c.tail = newNode
		return
	}

	// If the key exists, update the value and move to front.
	// The front of the list is the most recently used node.
	if node, ok := c.cache[key]; ok {
		node.value = value
		c.moveToFront(node)
		return
	}

	// If the key does not exist, create a new node
	newNode := &Node{
		key:   key,
		value: value,
	}
	c.cache[key] = newNode

	newNode.prev = c.head
	if oldHead, ok := c.cache[c.head.key]; ok {
		oldHead.next = newNode
	}
	c.head = newNode

	if c.tail == nil {
		c.tail = newNode
	}

	c.trimTail()
}

// Tr`im the tail of the list if it is over capacity.
func (c *LRUCache) trimTail() {
	if len(c.cache) > c.capacity {
		delete(c.cache, c.tail.key)
		c.tail = c.tail.next
	}
}

func (c *LRUCache) get(key string) int {

	if node, ok := c.cache[key]; ok {
		c.moveToFront(node)
		return node.value
	}

	return -1
}

func (c *LRUCache) moveToFront(node *Node) {
	if node == c.head || len(c.cache) == 1 {
		return
	}

	// Remove node from current position
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	node.next = nil
	node.prev = c.head
	c.head.next = node
	c.head = node

	// reset tail if needed
	if c.tail == nil {
		c.tail = node
	}

	if c.tail.prev == nil {
		return
	}

	// move c.tail to the last node
	for {
		if c.tail.prev == nil {
			break
		}
		c.tail = c.tail.prev
	}
}
