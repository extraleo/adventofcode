package main

import "slices"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	slices.Sort(coins)
	mod := amount % coins[len(coins)-1]
	if mod < coins[0]{
		return -1
	}
	if slices.Contains(coins, mod){

	}
}