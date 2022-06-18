//Design a data structure that follows the constraints of a Least Recently Used
//(LRU) cache.
//
// Implement the LRUCache class:
//
//
// LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
//
// int get(int key) Return the value of the key if the key exists, otherwise
//return -1.
// void put(int key, int value) Update the value of the key if the key exists.
//Otherwise, add the key-value pair to the cache. If the number of keys exceeds
//the capacity from this operation, evict the least recently used key.
//
//
// The functions get and put must each run in O(1) average time complexity.
//
//
// Example 1:
//
//
//Input
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//Output
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//Explanation
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // cache is {1=1}
//lRUCache.put(2, 2); // cache is {1=1, 2=2}
//lRUCache.get(1);    // return 1
//lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
//lRUCache.get(2);    // returns -1 (not found)
//lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
//lRUCache.get(1);    // return -1 (not found)
//lRUCache.get(3);    // return 3
//lRUCache.get(4);    // return 4
//
//
//
// Constraints:
//
//
// 1 <= capacity <= 3000
// 0 <= key <= 10⁴
// 0 <= value <= 10⁵
// At most 2 * 10⁵ calls will be made to get and put.
//
// Related Topics 设计 哈希表 链表 双向链表 👍 2080 👎 0

package q0146

//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	capacity int
	lruMap   map[int]*ListNode
	lruList  *LinkedList
}

type LinkedList struct {
	size int
	head *ListNode
	tail *ListNode
}

type ListNode struct {
	key   int
	value int
	prev  *ListNode
	next  *ListNode
}

// nil<-head(dummy) <-> 1 <-> 2 <-> tail(dummy)->nil

func (l *LinkedList) AddFirst(node *ListNode) {
	if node == nil {
		return
	}
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
	l.size++
}

func (l *LinkedList) Remove(node *ListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}

func (l *LinkedList) MoveToFirst(node *ListNode) {
	l.Remove(node)
	l.AddFirst(node)
}

func (l *LinkedList) RemoveLast() *ListNode {
	last := l.tail.prev
	l.Remove(last)
	return last
}

func (l *LinkedList) Size() int {
	return l.size
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		lruMap: make(map[int]*ListNode),
		lruList: &LinkedList{
			head: &ListNode{},
			tail: &ListNode{},
		},
		capacity: capacity,
	}
	cache.lruList.head.next = cache.lruList.tail
	cache.lruList.tail.prev = cache.lruList.head
	return cache
}

func (this LRUCache) Get(key int) int {
	v, ok := this.lruMap[key]
	if ok {
		this.lruList.MoveToFirst(v)
		return v.value
	} else {
		return -1
	}
}

func (this LRUCache) Put(key int, value int) {
	v, ok := this.lruMap[key]
	if ok {
		v.value = value
		this.lruList.MoveToFirst(v)
	} else {
		node := &ListNode{
			key:   key,
			value: value,
		}
		this.lruMap[key] = node
		this.lruList.AddFirst(node)
		if this.lruList.size > this.capacity {
			last := this.lruList.RemoveLast()
			delete(this.lruMap, last.key)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
