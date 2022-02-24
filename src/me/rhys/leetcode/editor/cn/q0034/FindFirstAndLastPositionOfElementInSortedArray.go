//Given an array of integers nums sorted in non-decreasing order, find the
//starting and ending position of a given target value.
//
// If target is not found in the array, return [-1, -1].
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Example 1:
// Input: nums = [5,7,7,8,8,10], target = 4
//Output: [3,4]
// Example 2:
// Input: nums = [5,7,7,8,8,10], target = 6
//Output: [-1,-1]
// Example 3:
// Input: nums = [], target = 0
//Output: [-1,-1]
//
//
// Constraints:
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// nums is a non-decreasing array.
// -10⁹ <= target <= 10⁹
//
// Related Topics 数组 二分查找 👍 1469 👎 0

package q0034

//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	var l, r int
	// 二分搜索边界，先搜左边界，再搜右边界
	left, right := 0, len(nums)

	// 左边界，就是右侧向左探测
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	// 补丁：
	// 首先，退出时left==right
	// target小于所有元素时，left等于0
	// target大于所有元素时，left等于len(nums)
	// 退出时，因为right==mid,right==left，所以left还可能是target，所以要判断一下nums[left]等不等于target，但要排除left=len(nums)的场景（越界了）
	if left == len(nums) || nums[left] != target {
		l = -1
	} else {
		l = left
	}
	// 右边界，就是左侧向右探测
	left, right = 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	// 补丁：
	// 首先退出时left=right，用谁都一样
	// target小于所有元素时，left（和right）会是0
	// target大于所有元素时，left会是len(nums)
	// 退出时，因为left=mid+1，所以nums[left]不等于target了（要还等于肯定还有下一个循环），nums[left-1]可能是，所以要判断nums[left-1]等不等于target，此时还要特殊处理left==0的场景（left-1越界了）
	if left == 0 || nums[left-1] != target {
		r = -1
	} else {
		r = left - 1
	}
	return []int{l, r}
}

//leetcode submit region end(Prohibit modification and deletion)
