/*
 * @lc app=leetcode.cn id=1419 lang=golang
 *
 * [1419] 数青蛙
 */

// @lc code=start
func minNumberOfFrogs(croakOfFrogs string) int {
	croak:=make(map[roue]int, 0,5)
	var count int
	for index, c:=range croakOfFrogs{
		croak[c]++
		if (index == 4){
			count=croak['c']
		}
	}

	if (croak['c'] == croak['r'] && croak['r'] == croak ['o'] && croak['o']==croak['a'] && croak['a'] ==croak['k'] ){
		return count
	}
	return -1
}
// @lc code=end

