package main
//给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充
//。
// 
// 
// 
//
// 示例 1： 
//
// 
//输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X"
//,"X"]]
//输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
//解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都
//会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
// 
//
// 示例 2： 
//
// 
//输入：board = [["X"]]
//输出：[["X"]]
// 
//
// 
//
// 提示： 
//
// 
// m == board.length 
// n == board[i].length 
// 1 <= m, n <= 200 
// board[i][j] 为 'X' 或 'O' 
// 
// 
// 
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 
// 👍 561 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func solve(board [][]byte)  {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows := len(board)
	cols := len(board[0])
	dx, dy := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}

	var dfs func([][]byte, int, int)
	dfs = func(board [][]byte, x int, y int) {
		if x < 0 || y < 0 || x >= rows || y >= cols || board[x][y] != 'O' {
			return
		}
		board[x][y] = '#'

		for i := 0; i < 4; i++ {
			nx, ny := x + dx[i], y + dy[i]
			dfs(board, nx, ny)
		}
	}

	for i := 0; i < rows; i++ {
		dfs(board, i, 0)
		dfs(board, i, cols - 1)
	}
	for i := 0; i < cols; i++ {
		dfs(board, 0, i)
		dfs(board, rows - 1, i)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
