//You are given an integer array coins representing coins of different
//denominations and an integer amount representing a total amount of money.
//
// Return the fewest number of coins that you need to make up that amount. If
//that amount of money cannot be made up by any combination of the coins, return -1.
//
//
// You may assume that you have an infinite number of each kind of coin.
//
//
// Example 1:
//
//
//Input: coins = [1,2,5], amount = 11
//Output: 3
//Explanation: 11 = 5 + 5 + 1
//
//
// Example 2:
//
//
//Input: coins = [2], amount = 3
//Output: -1
//
//
// Example 3:
//
//
//Input: coins = [1], amount = 0
//Output: 0
//
//
//
// Constraints:
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2³¹ - 1
// 0 <= amount <= 10⁴
//
// Related Topics 广度优先搜索 数组 动态规划 👍 1808 👎 0

package q0322

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	// 方案一：自顶向下迭代
	//memo := make([]int, amount+1)
	//for i := 0; i < amount+1; i++ {
	//	memo[i] = math.MaxInt
	//}
	//return subCoinChange(memo, coins, amount)

	// 方案二：自底向上循环 dp数组也是数组
	// dp[i]表示金额为i时，需要用dp[i]个金币，其实和memo是相同的
	// 循环就要从base case循环了，也就是dp[0]=0，amount等于0时，需要0个金币；然后循环这个数组
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		// 从金币里计算可以生成改金额的金币
		for _, c := range coins {
			if i-c < 0 {
				// 如果金额为负，则不用计算了，说明没有这样的金额
				continue
			}
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func subCoinChange(memo []int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if memo[amount] != math.MaxInt {
		return memo[amount]
	} else {
		res := amount + 1
		for _, v := range coins {
			subCount := subCoinChange(memo, coins, amount-v)
			if subCount == -1 {
				continue
			}
			res = min(subCount+1, res)
		}
		if res == amount+1 {
			memo[amount] = -1
		} else {
			memo[amount] = res
		}
	}
	return memo[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
