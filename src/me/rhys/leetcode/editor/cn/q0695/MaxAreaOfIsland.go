//You are given an m x n binary matrix grid. An island is a group of 1's (
//representing land) connected 4-directionally (horizontal or vertical.) You may assume
//all four edges of the grid are surrounded by water.
//
// The area of an island is the number of cells with a value 1 in the island.
//
// Return the maximum area of an island in grid. If there is no island, return 0
//.
//
//
// Example 1:
//
//
//Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,
//0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,
//0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]
//]
//Output: 6
//Explanation: The answer is not 11, because the island must be connected 4-
//directionally.
//
//
// Example 2:
//
//
//Input: grid = [[0,0,0,0,0,0,0,0]]
//Output: 0
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 50
// grid[i][j] is either 0 or 1.
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 796 👎 0

package q0695

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
/**
  TODO: 695
*/
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := flood(grid, i, j)
				maxArea = int(math.Max(float64(area), float64(maxArea)))
			}

		}

	}
	return maxArea
}

func flood(grid [][]int, i int, j int) int {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	// 淹没
	grid[i][j] = 0
	return 1 +
		flood(grid, i-1, j) +
		flood(grid, i+1, j) +
		flood(grid, i, j-1) +
		flood(grid, i, j+1)

}

//leetcode submit region end(Prohibit modification and deletion)
