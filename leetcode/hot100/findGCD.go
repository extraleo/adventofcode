package main

import "slices"

func findGCD(nums []int) int{
	slices.Sort(nums)
	return gcd(nums[len(nums)-1], nums[0])
}


func lcm(a, b int) int{
	return a/gcd(a, b) * b
}

func gcd(a, b int) int{
	for b!=0{
		mod := a%b
		a = b
		b = mod
	}
	return a
}
