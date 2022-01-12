package main

// given an array of coins, find the minimum number of coins needed to make change for a given amount
// coins = [1, 2, 5]
// amount = 11
// return 3 (11 = 5 + 5 + 1)
// coins = [2]
// amount = 3
// return -1 (no solution)
// coins = [1]
// amount = 0
// return 0 (0 = 0)
// coins = [1]
// amount = 1
// return 1 (1 = 1)
// coins = [1, 2, 5]
// amount = 3
// return 2 (3 = 1 + 1)
// coins = [1, 2, 5]
// amount = 4
// return 2 (4 = 2 + 1)
// coins = [1, 2, 5]
// amount = 5
// return 1 (5 = 5)
// coins = [1, 2, 5]
// amount = 6
// return 2 (6 = 5 + 1)
// coins = [1, 2, 5]
// amount = 7
// return 2 (7 = 5 + 2)
// coins = [1, 2, 5]
// amount = 8
// return 3 (8 = 5 + 5)
// coins = [1, 2, 5]
// amount = 9
// return 3 (9 = 5 + 5 + 1)
// coins = [1, 2, 5]
// amount = 10
// return 2 (10 = 5 + 5)
// coins = [1, 2, 5]
// amount = 11
// return 3 (11 = 5 + 5 + 1)
// coins = [1, 2, 5]
// amount = 12
// return 3 (12 = 5 + 5 + 2)
// coins = [1, 2, 5]
// amount = 13
// return 4 (13 = 5 + 5 + 5)
// coins = [1, 2, 5]
// amount = 14
// return 4 (14 = 5 + 5 + 5 + 1)
// coins = [1, 2, 5]
// amount = 15
// return 4 (15 = 5 + 5 + 5 + 2)
// coins = [1, 2, 5]
// amount = 16
// return 5 (16 = 5 + 5 + 5 + 5)
// coins = [1, 2, 5]
// amount = 17
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

