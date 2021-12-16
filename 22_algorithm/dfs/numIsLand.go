package main

import "fmt"

func numIslands(grid [][]byte) int {
	res := 0
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		// 检测是否超出边界
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}
		// 检测是否为海水
		if grid[i][j] == byte(0) {
			return
		}

		// 沉没陆地为海水
		grid[i][j] = byte(0)
		// 沉没陆地四周，可以防止走回头路
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == byte(1) {
				res++
				dfs(grid, i, j)
			}

		}
	}
	return res
}

func main() {
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', 0},
		[]byte{1, 1, 0, 1, 0},
		[]byte{1, 1, 0, 0, 0},
		[]byte{0, 0, 0, 0, 0},
	}
	fmt.Println(numIslands(grid), '1' == byte('1'))
}
