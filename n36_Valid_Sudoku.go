package main

func isValidSudoku(board [][]byte) bool {
	rowSeen := [9][9]bool{} //检查行
	colSeen := [9][9]bool{} //检查列
	boxSeen := [9][9]bool{} //检查小方格

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			x := board[i][j] - '1'
			k := (i/3)*3 + j/3 //将i,j坐标映射到3 * 3的小方格上，用来表示是第几个方格（从左到右，从上到下）
			if rowSeen[i][x] || colSeen[j][x] || boxSeen[k][x] == true {
				return false
			}
			rowSeen[i][x] = true
			colSeen[j][x] = true
			boxSeen[k][x] = true
		}
	}
	return true
}
