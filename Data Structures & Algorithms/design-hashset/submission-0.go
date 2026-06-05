type hashNode struct {
	key   int
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

func Constructor() MyHashSet{
	return MyHashSet{nodeList: make([]hashNode, 1_000_000)}

}

func (this *MyHashSet) Add(key int) {
	hash := this.hashKey(key)
	location := &this.nodeList[hash]

	// find the right node by going to the end of the linked list
	for location.next != nil {
		if location.next.key == key {
			return
		}
		location = location.next
	}
	location.next = &hashNode{key: key}
}

func (this *MyHashSet) Remove(key int) {
	hash := this.hashKey(key)

	location := &this.nodeList[hash]

	for location.next != nil {
		if location.next.key == key {
			location.next = location.next.next
			break
		}
		location = location.next
	}

}

func (this *MyHashSet) Contains(key int) bool {
	hash := this.hashKey(key)


	location := this.nodeList[hash]

	for location.next != nil {
		if location.next.key == key {
			return true
		}
		location = *location.next
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 