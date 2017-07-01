package LRUCache

type LinkNode struct {
	Key   int
	Value int
	Prev  *LinkNode
	Next  *LinkNode
}

type LRUCache struct {
	Cache    map[int]*LinkNode
	Head     *LinkNode
	Tail     *LinkNode
	Count    int
	Capacity int
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{}
	tail := &LinkNode{}

	head.Next = tail
	tail.Prev = head

	return LRUCache{
		Cache:    map[int]*LinkNode{},
		Head:     head,
		Tail:     tail,
		Count:    0,
		Capacity: capacity,
	}
}

func (l *LRUCache) addNode(node *LinkNode) {
	node.Prev = l.Head
	node.Next = l.Head.Next

	l.Head.Next.Prev = node
	l.Head.Next = node
}

func (l *LRUCache) removeNode(node *LinkNode) {
	pre := node.Prev
	next := node.Next

	pre.Next = next
	next.Prev = pre
}

func (l *LRUCache) moveToHead(node *LinkNode) {
	l.removeNode(node)
	l.addNode(node)
}

func (l *LRUCache) popTail() *LinkNode {
	tail := l.Tail.Prev
	l.removeNode(tail)

	return tail
}

func (l *LRUCache) Put(key int, value int) {
	node, ok := l.Cache[key]
	if !ok {
		node := &LinkNode{
			Key:   key,
			Value: value,
		}

		l.Cache[key] = node
		l.addNode(node)

		l.Count = l.Count + 1
		if l.Count > l.Capacity {
			detail := l.popTail()
			delete(l.Cache, detail.Key)
			l.Count = l.Count - 1
		}
	} else {
		node.Value = value
		l.moveToHead(node)
	}
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.Cache[key]
	if !ok {
		return -1
	}
	l.moveToHead(node)
	return node.Value
}

func (l *LRUCache) GetNodeCount() int {
	return l.Count
}

func (l *LRUCache) GetCapacity() int {
	return l.Capacity
}
