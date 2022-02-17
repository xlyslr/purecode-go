//There is a car with capacity empty seats. The vehicle only drives east (i.e.,
//it cannot turn around and drive west).
//
// You are given the integer capacity and an array trips where trips[i] = [
//numPassengersi, fromi, toi] indicates that the iᵗʰ trip has numPassengersi
//passengers and the locations to pick them up and drop them off are fromi and toi
//respectively. The locations are given as the number of kilometers due east from the
//car's initial location.
//
// Return true if it is possible to pick up and drop off all passengers for all
//the given trips, or false otherwise.
//
//
// Example 1:
//
//
//Input: trips = [[2,1,5],[3,3,7]], capacity = 4
//Output: false
//
//
// Example 2:
//
//
//Input: trips = [[2,1,5],[3,3,7]], capacity = 5
//Output: true
//
//
//
// Constraints:
//
//
// 1 <= trips.length <= 1000
// trips[i].length == 3
// 1 <= numPassengersi <= 100
// 0 <= fromi < toi <= 1000
// 1 <= capacity <= 10⁵
//
// Related Topics 数组 前缀和 排序 模拟 堆（优先队列） 👍 147 👎 0

package q1094

//leetcode submit region begin(Prohibit modification and deletion)

func carPooling(trips [][]int, capacity int) bool {
	// 题中说小于等于1000，因为初始都是0，所以差分数组不需要初始化
	nums := make([]int, 1001)

	for i := 0; i < len(trips); i++ {
		incr(nums, trips[i][0], trips[i][1], trips[i][2])
	}
	result := result(nums)
	for i := 0; i < len(result); i++ {
		if capacity < result[i] {
			return false
		}
	}
	return true
}

func incr(nums []int, num, from, to int) {
	nums[from] += num
	nums[to] -= num
}

func result(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		result[i] = nums[i] + result[i-1]
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
