//Given an m x n 2D binary grid grid which represents a map of '1's (land) and
//'0's (water), return the number of islands.
//
// An island is surrounded by water and is formed by connecting adjacent lands
//horizontally or vertically. You may assume all four edges of the grid are all
//surrounded by water.
//
//
// Example 1:
//
//
//Input: grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//Output: 1
//
//
// Example 2:
//
//
//Input: grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//Output: 3
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] is '0' or '1'.
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 1764 👎 0

package q0200

//leetcode submit region begin(Prohibit modification and deletion)
/*
思路：https://labuladong.github.io/algo/4/29/106/
分解子问题（想想和遍历二叉树没啥区别，多遍历几个方向）：遍历二维数组，当遇到土地则加1，然后顺便把这个岛淹掉，继续遍历
*/
func numIslands(grid [][]byte) int {
	num := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				num++
				flood(grid, i, j)
			}
		}
	}
	return num
}

func flood(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == 1 {
		grid[i][j] = 0
	} else {
		return
	}
	flood(grid, i-1, j)
	flood(grid, i+1, j)
	flood(grid, i, j-1)
	flood(grid, i, j+1)
}

//leetcode submit region end(Prohibit modification and deletion)
