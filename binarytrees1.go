package main

// given a non-empty binary tree T consisting of N nodes, return the length of longest zig-zag starting at root.
// N is an integer within the range [1, 100,000].
// the height of tree T is within range [1, 800].
// Definition for a binary tree node.
type Tree struct {
	Val   int
	L  *Tree
	R *Tree
}

// corner case:
func longestZigZag(root *Tree) int {
	if root == nil {
		return 0
	}
	var longest int
	var dfs func(node *Tree, isLeft bool)
	dfs = func(node *Tree, isLeft bool) {
		if node.L == nil && node.R == nil {
			if isLeft {
				longest = max(longest, 1)
			} else {
				longest = max(longest, 2)
			}
			return
		}
		// handle corner case:
		if node.L == nil {
			dfs(node.R, !isLeft)
			return
		}
		if node.R == nil {
			dfs(node.L, !isLeft)
			return
		}
		if isLeft {
			dfs(node.L, !isLeft)
			dfs(node.R, !isLeft)
		} else {
			dfs(node.R, !isLeft)
			dfs(node.L, !isLeft)
		}
	}
	dfs(root, true)
	return longest
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}