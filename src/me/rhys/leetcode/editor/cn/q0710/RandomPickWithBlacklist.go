//You are given an integer n and an array of unique integers blacklist. Design
//an algorithm to pick a random integer in the range [0, n - 1] that is not in
//blacklist. Any integer that is in the mentioned range and not in blacklist should
//be equally likely to be returned.
//
// Optimize your algorithm such that it minimizes the number of calls to the
//built-in random function of your language.
//
// Implement the Solution class:
//
//
// Solution(int n, int[] blacklist) Initializes the object with the integer n
//and the blacklisted integers blacklist.
// int pick() Returns a random integer in the range [0, n - 1] and not in
//blacklist.
//
//
//
// Example 1:
//
//
//Input
//["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
//[[7, [2, 3, 5]], [], [], [], [], [], [], []]
//Output
//[null, 0, 4, 1, 6, 1, 0, 4]
//
//Explanation
//Solution solution = new Solution(7, [2, 3, 5]);
//solution.pick(); // return 0, any integer from [0,1,4,6] should be ok. Note
//that for every call of pick,
//                 // 0, 1, 4, and 6 must be equally likely to be returned (i.e.
//, with probability 1/4).
//solution.pick(); // return 4
//solution.pick(); // return 1
//solution.pick(); // return 6
//solution.pick(); // return 1
//solution.pick(); // return 0
//solution.pick(); // return 4
//
//
//
// Constraints:
//
//
// 1 <= n <= 10⁹
// 0 <= blacklist.length <- min(10⁵, n - 1)
// 0 <= blacklist[i] < n
// All the values of blacklist are unique.
// At most 2 * 10⁴ calls will be made to pick.
//
// Related Topics 哈希表 数学 二分查找 排序 随机化 👍 88 👎 0

package q0710

import "math/rand"

//leetcode submit region begin(Prohibit modification and deletion)

// Solution 题意是一个连续的数组中，要剔除几个黑名单数据，然后随机返回。尽量少调用rand，也就是要求还在一个连续的范围内，一次性返回
// 所以：把黑名单元素用数组最后的元素替代，这样一个新的"全白数组"就可以随机返回了
// 如果保存数组的话，数组过大会内存溢出，所以值保存黑名单到白数据即可
type Solution struct {
	b2w map[int]int // 黑到白的映射
	sz  int         // 全白数组的大小 n- len(blacklist)
}

func Constructor(n int, blacklist []int) Solution {
	b2w := map[int]int{}
	sz := n - len(blacklist)
	blkMap := map[int]int{}
	for _, v := range blacklist {
		blkMap[v] = 666
	}
	// 给在sz以内的黑数据寻找sz以外的白数据
	last := n - 1 // 这就是从后边找的白数据
	for _, v := range blacklist {
		// 如果本身黑数据就在sz以外就不用放进索引的了
		if v >= sz {
			continue
		}

		_, isBlk := blkMap[last]
		for isBlk {
			last--
			_, isBlk = blkMap[last]
		}
		b2w[v] = last
		last--
	}

	return Solution{b2w, sz}
}

func (this *Solution) Pick() int {
	r := rand.Intn(this.sz)
	if v, ok := this.b2w[r]; ok {
		return v
	} else {
		return r
	}

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
//leetcode submit region end(Prohibit modification and deletion)
