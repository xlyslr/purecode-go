//Given an array of integers nums, sort the array in ascending order.
//
//
// Example 1:
// Input: nums = [5,2,3,1]
//Output: [1,2,3,5]
// Example 2:
// Input: nums = [5,1,1,2,0,0]
//Output: [0,0,1,1,2,5]
//
//
// Constraints:
//
//
// 1 <= nums.length <= 5 * 10⁴
// -5 * 10⁴ <= nums[i] <= 5 * 10⁴
//
// Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 502 👎 0

package q0912

//leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	//for i := 0; i < len(nums)-1; i++ {
	//	for j := 0; j < len(nums)-i-1; j++ {
	//		if nums[j] > nums[j+1] {
	//			nums[j], nums[j+1] = nums[j+1], nums[j]
	//		}
	//	}
	//}
	//return nums
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(nums, lo, hi)
	quickSort(nums, lo, p-1)
	quickSort(nums, p+1, hi)
}

func partition(nums []int, lo, hi int) int {
	// 目的是使得左边都小于基准值，右边都大于基准值
	pivot := nums[lo]
	i := lo + 1
	j := hi
	for i <= j {
		// 从左向右找，第一个大于基准数的元素
		for i < hi && nums[i] <= pivot {
			i++
		}
		// 从右向做找，第一个小于等于基准数的元素
		for j > lo && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		// 然后交换
		swap(nums, i, j)
	}
	// 为什么和j交换，因为j是小于等于基准数的元素，而lo肯定在左边，所以用他来交换
	swap(nums, lo, j)
	return j
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]

}

//leetcode submit region end(Prohibit modification and deletion)
