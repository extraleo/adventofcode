package utils


func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	res := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		res = LCM(res, integers[i])
	}
	return res
}


func Max(item ...int) int {
	max := 0
	for _, i := range item {
		if i==0||i > max {
			max = i
		}
	}
	return max
}

func Min(item ...int) int {
	min := 0
	for _, i := range item {
		if i==0||i < min {
			min = i
		}
	}
	return min
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}