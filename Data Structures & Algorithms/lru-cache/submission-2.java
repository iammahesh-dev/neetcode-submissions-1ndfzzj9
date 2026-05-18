class LRUCache {
	private class Node {
		int key;
		int value;
		Node next;
		Node prev;
		Node(int key, int value){
			this.key = key;
			this.value = value;
			this.next = null;
			this.prev = null;
		}
	}

	private int capacity;
	private HashMap<Integer, Node> cache;
	private Node head;
	private Node tail;

    public LRUCache(int capacity) {
        this.capacity = capacity;
		this.cache = new HashMap<>();
		this.head = new Node(0,0);
		this.tail = new Node(0,0);
		this.head.next = this.tail;
		this.tail.prev = this.head;
    }

	private void add(Node node){
		Node prev = this.head.next;
		node.next = prev;
		node.prev = this.head;
		this.head.next = node;
		prev.prev = node;
	}

	private void remove(Node node){
		Node nxt = node.next;
		Node prev = node.prev;
		prev.next = nxt;
		nxt.prev = prev;
	}
    
    public int get(int key) {
        if(cache.containsKey(key)){
			Node node = cache.get(key);
			remove(node);
			add(node);
			return node.value;
		}
		return -1;
    }
    
    public void put(int key, int value) {
        if(cache.containsKey(key)){
			remove(cache.get(key));
		}

		Node node = new Node(key, value);
		cache.put(key, node);
		add(node);

		if(cache.size() > capacity){
			Node last = this.tail.prev;
			remove(last);
			cache.remove(last.key);
		}

    }
}
