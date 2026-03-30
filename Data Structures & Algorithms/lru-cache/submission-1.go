type Node struct{
	key, val int
	prev, next *Node
}
type LRUCache struct {
    cap int
	cache map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
    return LRUCache{
		cap: capacity,
		cache: make(map[int]*Node),
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) add(key, value int) {
	node := &Node{key: key, val: value}
	node.next = this.head.next
    node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
	this.cache[key] = node
}

func (this *LRUCache) remove(key int){
	node, _ := this.cache[key]
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
	delete(this.cache, key)
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
		this.remove(key)
		this.add(key, node.val)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    fmt.Println(this.cache)
    if _, ok := this.cache[key]; ok {
		this.remove(key)
	}
	this.add(key, value)
	if len(this.cache) > this.cap {
		this.remove(this.tail.prev.key)
	}
}
