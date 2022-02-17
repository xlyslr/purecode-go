//There are n flights that are labeled from 1 to n.
//
// You are given an array of flight bookings bookings, where bookings[i] = [
//firsti, lasti, seatsi] represents a booking for flights firsti through lasti (
//inclusive) with seatsi seats reserved for each flight in the range.
//
// Return an array answer of length n, where answer[i] is the total number of
//seats reserved for flight i.
//
//
// Example 1:
//
//
//Input: bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
//Output: [10,55,45,25,25]
//Explanation:
//Flight labels:        1   2   3   4   5
//Booking 1 reserved:  10  10
//Booking 2 reserved:      20  20
//Booking 3 reserved:      25  25  25  25
//Total seats:         10  55  45  25  25
//Hence, answer = [10,55,45,25,25]
//
//
// Example 2:
//
//
//Input: bookings = [[1,2,10],[2,2,15]], n = 2
//Output: [10,25]
//Explanation:
//Flight labels:        1   2
//Booking 1 reserved:  10  10
//Booking 2 reserved:      15
//Total seats:         10  25
//Hence, answer = [10,25]
//
//
//
//
// Constraints:
//
//
// 1 <= n <= 2 * 10⁴
// 1 <= bookings.length <= 2 * 10⁴
// bookings[i].length == 3
// 1 <= firsti <= lasti <= n
// 1 <= seatsi <= 10⁴
//
// Related Topics 数组 前缀和 👍 308 👎 0

package q1109

//leetcode submit region begin(Prohibit modification and deletion)
func corpFlightBookings(bookings [][]int, n int) []int {
	// 差分数组
	labels := make([]int, n+1)
	for i := 0; i < len(bookings); i++ {
		labels[bookings[i][0]] += bookings[i][2]
		// 为什么加1？因为last包含在加的范围内
		last := bookings[i][1] + 1
		// 为什么n+1？因为差分数组的0位忽略了，数组长度是n+1
		if last < n+1 {
			labels[last] -= bookings[i][2]
		}
	}
	seats := make([]int, n+1)
	seats[0] = labels[0]
	for i := 1; i < len(labels); i++ {
		seats[i] = labels[i] + seats[i-1]
	}
	return seats[1:]
}

//leetcode submit region end(Prohibit modification and deletion)
