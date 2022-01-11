package backtrack

func combinationSumBacktrack(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	backtrackCombinationSum(candidates, target, 0, path, &res)
	return res
}

func backtrackCombinationSum(candidates []int, target int, start int, path []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			return
		}
		path = append(path, candidates[i])
		backtrackCombinationSum(candidates, target-candidates[i], i, path, res)
		path = path[:len(path)-1]
	}
}