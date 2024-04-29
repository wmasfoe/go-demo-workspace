package leetcode0003

func lengthOfLongestSubstring(s string) int {
	rangeedMap := make(map[byte]bool)
	index := 0
	res := 0

	for _, v := range []byte(s) {
		if rangeedMap[v] {
			if index > res {
				res = index
			}

			rangeedMap = make(map[byte]bool)
			// 清空之后添加当前的 byte
			rangeedMap[v] = true

			index = 1
		} else {
			rangeedMap[v] = true
			index++
		}
	}

	if index > res {
		res = index
	}

	return res
}

func RunDemo() {
	//lengthOfLongestSubstring("bbbbb")
	lengthOfLongestSubstring("dvdf")
}
