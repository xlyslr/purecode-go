//Given a 2D matrix matrix, handle multiple queries of the following type:
//
//
// Calculate the sum of the elements of matrix inside the rectangle defined by
//its upper left corner (row1, col1) and lower right corner (row2, col2).
//
//
// Implement the NumMatrix class:
//
//
// NumMatrix(int[][] matrix) Initializes the object with the integer matrix
//matrix.
// int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the
//elements of matrix inside the rectangle defined by its upper left corner (row1,
//col1) and lower right corner (row2, col2).
//
//
//
// Example 1:
//
//
//Input
//["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
//[[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0,
//3, 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]
//Output
//[null, 8, 11, 12]
//
//Explanation
//NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2,
//0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
//numMatrix.sumRegion(2, 1, 4, 3); // return 8 (i.e sum of the red rectangle)
//numMatrix.sumRegion(1, 1, 2, 2); // return 11 (i.e sum of the green rectangle)
//
//numMatrix.sumRegion(1, 2, 2, 4); // return 12 (i.e sum of the blue rectangle)
//
//
//
// Constraints:
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 200
// -10⁵ <= matrix[i][j] <= 10⁵
// 0 <= row1 <= row2 < m
// 0 <= col1 <= col2 < n
// At most 10⁴ calls will be made to sumRegion.
//
// Related Topics 设计 数组 矩阵 前缀和 👍 348 👎 0

package q0304

//leetcode submit region begin(Prohibit modification and deletion)
type NumMatrix struct {
	preSums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var row, col int
	if len(matrix) == 0 {
		row, col = 1, 1
	} else {
		row, col = len(matrix)+1, len(matrix[0])+1
	}
	preSums := make([][]int, row)
	for i, _ := range preSums {
		preSums[i] = make([]int, col)
	}
	// 数组的默认值为0，第0列和第0行不参与，
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			preSums[i][j] = preSums[i-1][j] + preSums[i][j-1] - preSums[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{
		preSums: preSums,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 只有大的矩阵点加1，为什么？因为小的矩阵位是包含在和之中的，实际是要减1的，只不过preSums本身已经加1了。
	// 大的矩阵点比如是3,3，实际3,3到0,0的和是4x4个格子的和，保存在preSums[4,4]中
	return this.preSums[row2+1][col2+1] - this.preSums[row1][col2+1] - this.preSums[row2+1][col1] + this.preSums[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
//leetcode submit region end(Prohibit modification and deletion)
