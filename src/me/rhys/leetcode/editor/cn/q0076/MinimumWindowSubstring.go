//Given two strings s and t of lengths m and n respectively, return the minimum
//window substring of s such that every character in t (including duplicates) is
//included in the window. If there is no such substring, return the empty string
//"".
//
// The testcases will be generated such that the answer is unique.
//
// A substring is a contiguous sequence of characters within the string.
//
//
// Example 1:
//
//
//Input: s = "ADOBECODEBANC", t = "ABC"
//Output: "BANC"
//Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C'
//from string t.
//
//
// Example 2:
//
//
//Input: s = "a", t = "a"
//Output: "a"
//Explanation: The entire string s is the minimum window.
//
//
// Example 3:
//
//
//Input: s = "a", t = "aa"
//Output: ""
//Explanation: Both 'a's from t must be included in the window.
//Since the largest window of s only has one 'a', return empty string.
//
//
//
// Constraints:
//
//
// m == s.length
// n == t.length
// 1 <= m, n <= 10⁵
// s and t consist of uppercase and lowercase English letters.
//
//
//
//Follow up: Could you find an algorithm that runs in O(m + n) time? Related
//Topics 哈希表 字符串 滑动窗口 👍 1651 👎 0

package q0076

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	sa := []rune(s)
	ta := []rune(t)
	// need只是用来索引的，值是用来t字符串中重复字符的个数
	need := make(map[rune]int)
	for _, r := range ta {
		need[r]++
	}

	window := make(map[rune]int)

	left, right, valid := 0, 0, 0

	// 取最短的字符串，肯定有个值来更新
	resultLen := len(s) + 1 // 这个初始值没有实际意义，只不过为了后边判断时，不用再单独判断是不是初始值；这个初始值肯定是覆盖right-left<resultLen的场景的
	result := ""
	for right < len(s) {
		// 先向右扩展窗口，以找到符合要求的子字符串
		r := sa[right]
		right++

		// 如果是t中所包含的字符，并且字符出现的个数与t中相符，则说明符合要求
		// 这里 need就是t的额外索引，只读的；
		// valid 用来判断匹配字符的个数，当匹配字符的个数与need的长度相等了，则说明都匹配上了
		// window 是用来保存窗口内符合规则的字符的信息的，这是动态的
		if _, ok := need[r]; ok {
			window[r]++
			if window[r] == need[r] {
				valid++
			}
		}

		// 如果都找到了，则缩小左侧，以找到最小字符串
		for valid == len(need) {
			if right-left < resultLen {
				resultLen = right - left
				result = string(sa[left:right])
			}

			l := sa[left]
			left++

			if _, ok := need[l]; ok {
				if window[l] == need[l] {
					valid--
				}
				window[l]--
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
