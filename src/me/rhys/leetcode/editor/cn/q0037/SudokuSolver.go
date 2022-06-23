//Write a program to solve a Sudoku puzzle by filling the empty cells.
//
// A sudoku solution must satisfy all of the following rules:
//
//
// Each of the digits 1-9 must occur exactly once in each row.
// Each of the digits 1-9 must occur exactly once in each column.
// Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-
//boxes of the grid.
//
//
// The '.' character indicates empty cells.
//
//
// Example 1:
//
//
//Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5
//",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".
//",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".
//","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5
//"],[".",".",".",".","8",".",".","7","9"]]
//Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4
//","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3
//"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],[
//"9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3",
//"4","5","2","8","6","1","7","9"]]
//Explanation: The input board is shown above and the only valid solution is
//shown below:
//
//
//
//
//
// Constraints:
//
//
// board.length == 9
// board[i].length == 9
// board[i][j] is a digit or '.'.
// It is guaranteed that the input board has only one solution.
//
// Related Topics 数组 回溯 矩阵 👍 1305 👎 0

package q0037

//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	traverse(board, 0, 0)
}

func traverse(board [][]byte, row, col int) {
	if col == 9 { // 等于9时，已经过了，因为索引是从0开始的
		traverse(board, row+1, 0)
		return
	}
	if row == 9 { // 完成
		return
	}
	if board[row][col] != '.' { // 如果有数组直接跳过
		traverse(board, row, col+1)
		return
	}

	// 从1~9做选择
	var i byte
	for i = '1'; i <= '9'; i++ {
		if !isValid(board, row, col, i) {
			continue
		}
		board[row][col] = i
		traverse(board, row, col+1)
		board[row][col] = '.'
	}
}

// 三个条件
// 同一行不能重复
// 同一列不能重复
// 9个3*3不能重复
func isValid(board [][]byte, row, col int, val byte) bool {
	// 行不重复 & 列不重复
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
		if board[i][col] == val {
			return false
		}
		// 3*3不重复
		if board[(row/3)*3+i/3][(col/3)*3+i%3] == val {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
