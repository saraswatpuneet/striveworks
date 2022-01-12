package main

// given mxn grid, find the number of islands
// grid = [
// 	[1, 1, 0, 0, 0],
// 	[0, 1, 0, 0, 1],
// 	[1, 0, 0, 1, 1],
// 	[0, 0, 0, 0, 0],
// 	[1, 0, 1, 0, 1]
// ]
// return 1 (1 island)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}