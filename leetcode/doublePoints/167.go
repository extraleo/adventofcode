package main

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	result := []int{}
	for {
		if numbers[left]+numbers[right] == target {
			result = append(result, left+1)
			result = append(result, right+1)
			return result
		}

		if numbers[left]+numbers[right] > target {
			right--

		}
		if numbers[left]+numbers[right] < target {
			left++
		}
	}
}
