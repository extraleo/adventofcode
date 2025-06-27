package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)   
	ans := make([]int, n)
	st := []int{}
	for i:=n-1;i>=0;i--{
			t := temperatures[i]
			// 如果当前温度 >= 栈顶温度
			// remove 栈顶
			for len(st) > 0 && t >= temperatures[st[len(st)-1]]{
					st = st[:len(st)-1]
			}
			if len(st) > 0{
					ans[i] = st[len(st)-1] - i
			}
			st = append(st, i)
	}
	return ans
}