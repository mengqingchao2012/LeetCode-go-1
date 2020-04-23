package main

func exist(board [][]byte, word string) bool {
	row := len(board)
	if row == 0 { return false }
	col := len(board[0])
	if col == 0 { return false }

	// 构造一个辅助二维数组，用来标记检查过的位置
	visited := make([][]bool, row)
	for idx := range visited {
		visited[idx] = make([]bool, col)
	}

	// 遍历整个数组，进行查询
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if checkExist(board, visited, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func checkExist(board [][]byte, visited [][]bool, i, j int, word string, idx int) bool {
	// 退递归条件1：查找索引已经等于单词长度，说明已经找到了单词
	if idx == len(word) { return true }
	// 退递归条件2：i, j 越界，或者当前位置已经检查过（visited[i][j]=true)，或者当前位置的字符与需要查找的字符不匹配
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || visited[i][j] || board[i][j] != word[idx] {
		return false
	}

	// 开始进入回溯算法
	visited[i][j] = true // 先标记当前位置为已检查
	// 检查当前位置的上下左右四个方向，只要有任意一个方向符合要求，就返回 true
	exists := checkExist(board, visited, i + 1, j, word, idx + 1) ||
		      checkExist(board, visited, i - 1, j, word, idx + 1) ||
		      checkExist(board, visited, i, j + 1, word, idx + 1) ||
		      checkExist(board, visited, i, j - 1, word, idx + 1)
	visited[i][j] = false // 回溯，退递归
	return exists
}
