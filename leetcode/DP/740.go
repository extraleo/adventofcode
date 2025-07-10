package dp


// 看示例 2，nums=[2,2,3,3,3,4]。如果我们选了一个等于 3 的数，那么所有等于 2 和等于 4 的数都被删除，也就是都不能选。选了一个 3 后，剩下的 3 可以继续选。所以如果要选 3，所有的 3 都要选。

// 这种「相邻数字不能都选」联想到 198. 打家劫舍。

// 把 nums 转换成一个值域数组 a，其中 a[i] 表示 nums 中的等于 i 的元素之和。上面的例子中，a=[0,0,4,9,4]。因为 nums 中有 3 个 3，所以 a[3]=3+3+3=9。


func deleteAndEarn(nums []int) int {
	a := make([]int, maxEle(nums)+1)
	for _, x := range(nums){
			a[x] += x
	}
	return _rob(a)
}

func _rob(nums []int) int{
	f0, f1 := 0, 0
	for _, x := range(nums){
			f0, f1 = f1 , max(f1, f0 + x)
	}
	return f1
}

func maxEle(nums []int) int{
	ans := -1
	for _, x := range(nums){
			ans = max(ans,x)
	}
	return ans
}