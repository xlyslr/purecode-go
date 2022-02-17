//Given an integer array nums, handle multiple queries of the following type:
//
//
// Calculate the sum of the elements of nums between indices left and right
//inclusive where left <= right.
//
//
// Implement the NumArray class:
//
//
// NumArray(int[] nums) Initializes the object with the integer array nums.
// int sumRange(int left, int right) Returns the sum of the elements of nums
//between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... +
//nums[right]).
//
//
//
// Example 1:
//
//
//Input
//["NumArray", "sumRange", "sumRange", "sumRange"]
//[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
//Output
//[null, 1, -1, -3]
//
//Explanation
//NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
//numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
//numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
//numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -10⁵ <= nums[i] <= 10⁵
// 0 <= left <= right < nums.length
// At most 10⁴ calls will be made to sumRange.
//
// Related Topics 设计 数组 前缀和 👍 416 👎 0

package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	preSums []int
	//nums    []int
}

//func Constructor(nums []int) NumArray {
//	// 使用前缀和优化
//	// 为了避免索引地址的0特殊处理，所以0号索引不用，从1号索引开始用，在取得的时候通过+1取
//	preSums := make([]int, len(nums)+1)
//	for i := 1; i < len(preSums); i++ {
//		preSums[i] = preSums[i-1] + nums[i-1]
//	}
//	return NumArray{preSums: preSums}
//}
func (this *NumArray) SumRange(left int, right int) int {
	// right+1取当前0-right的和，left取left-1的和
	return this.preSums[right+1] - this.preSums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
