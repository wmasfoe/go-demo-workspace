package leetcode0705

type MyHashSet struct {
	val []int
}

func Constructor() MyHashSet {
	m := new(MyHashSet)
	return *m
}

func (m *MyHashSet) Add(key int) {
	isRepeat := false
	for _, v := range m.val {
		if key == v {
			isRepeat = true
			break
		}
	}

	if !isRepeat {
		m.val = append(m.val, key)
	}
}

func (m *MyHashSet) Remove(key int) {
	i := -1
	for index, v := range m.val {
		if v == key {
			i = index
		}
	}

	if i > -1 {
		m.val = append(m.val[0:i], m.val[i+1:]...)
	}
}

func (m *MyHashSet) Contains(key int) bool {
	hasVal := false
	for _, v := range m.val {
		if key == v {
			hasVal = true
			break
		}
	}

	return hasVal
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
