package main

func maxVowels(s string, k int) int {
	max := 0
	old := 0
	for i:=0; i< len(s); i++ {
		toAdd := 0
		toMin :=0
		in := i
		out := in - k
		
		if isVowels(s[in]){
			toAdd=1
		}

		if out >=0 && isVowels(s[out]){
			toMin = -1
		}
		
	old = old + toAdd + toMin
		max = maxMax(old, max)
		if max == k {
			return max
		}

	}
	return max
}

func isVowels(r byte) bool{
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}

func maxMax(a,b int) int{
	if a > b {
		return a

	}
	return b
}