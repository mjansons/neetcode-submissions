
type hashNode struct {
	key   int
	value int
	next  *hashNode
}

type MyHashSet struct {
	nodeList []hashNode
}

func (this *MyHashSet) hashKey(key int) int {
	hash := 0
	hash = (hash*31 + key) % len(this.nodeList)

	return hash
}

func Constructor() MyHashSet {
	return MyHashSet{nodeList: make([]hashNode, 1_000_000)}

}

func (this *MyHashSet) Put(key int, value int) {
	hash := this.hashKey(key)
	location := &this.nodeList[hash]

	// find the right node by going to the end of the linked list
	for location.next != nil {
		if location.next.key == key {
			location.next.value = value
			return
		}
		location = location.next
	}
	location.next = &hashNode{key: key, value: value}
}

func (this *MyHashSet) Remove(key int) {
	hash := this.hashKey(key)
	location := &this.nodeList[hash]

	for location.next != nil {
		if location.next.key == key {
			location.next = location.next.next
			return
		}
		location = location.next
	}

}

func (this *MyHashSet) Get(key int) int {
	hash := this.hashKey(key)

	location := this.nodeList[hash]

	for location.next != nil {
		if location.next.key == key {
			return location.next.value
		}
		location = *location.next
	}
	return -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */