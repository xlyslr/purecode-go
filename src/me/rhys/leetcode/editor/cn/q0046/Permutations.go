//Given an array nums of distinct integers, return all the possible
//permutations. You can return the answer in any order.
//
//
// Example 1:
// Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// Example 2:
// Input: nums = [0,1]
//Output: [[0,1],[1,0]]
// Example 3:
// Input: nums = [1]
//Output: [[1]]
//
//
// Constraints:
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// All the integers of nums are unique.
//
// Related Topics 数组 回溯 👍 2081 👎 0

package q0046

//leetcode submit region begin(Prohibit modification and deletion)
var used map[int]bool
var res [][]int

func permute(nums []int) [][]int {
	used = make(map[int]bool, len(nums))
	res = make([][]int, 0)
	var list []int
	traverse(nums, list)
	return res
}

func traverse(nums, list []int) {
	// 如果全是true则是一种排列
	if len(nums) == len(list) {
		l := make([]int, len(nums))
		copy(l, list)
		res = append(res, l)
		return
	}

	for _, num := range nums {
		if used[num] {
			continue
		}
		list = append(list, num)
		used[num] = true
		traverse(nums, list)
		used[num] = false
		list = list[:len(list)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
