//Given two strings s1 and s2, return true if s2 contains a permutation of s1,
//or false otherwise.
//
// In other words, return true if one of s1's permutations is the substring of
//s2.
//
//
// Example 1:
//
//
//Input: s1 = "ab", s2 = "eidbaooo"
//Output: true
//Explanation: s2 contains one permutation of s1 ("ba").
//
//
// Example 2:
//
//
//Input: s1 = "ab", s2 = "eidboaoo"
//Output: false
//
//
//
// Constraints:
//
//
// 1 <= s1.length, s2.length <= 10⁴
// s1 and s2 consist of lowercase English letters.
//
// Related Topics 哈希表 双指针 字符串 滑动窗口 👍 572 👎 0

package q0567

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, t string) bool {
	s := []rune(s1)

	need, window := make(map[rune]int), make(map[rune]int)
	for _, i2 := range t {
		need[i2]++
	}

	left, right, valid := 0, 0, 0

	// 右迁移循环
	for right < len(s) {
		r := s[right]
		right++
		// 右前移操作
		if _, ok := need[r]; ok {
			window[r]++
			if need[r] == window[r] {
				valid++
			}
		}

		// TODO for 左前移循环；不需要取最小，所以优化了
		// 问题是什么时候左迁移，在改题中，要保持长度和s2一致，所以要时刻迁移；在时刻迁移的情况下，只要发现valid相等，则就匹配上了
		for right-left >= len(t) {
			if valid == len(need) {
				return true
			}
			l := s[left]
			left++
			if _, ok := need[l]; ok {
				if window[l] == need[l] {
					valid--
				}
				window[l]--
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
