package main


func carPooling(trips [][]int, capacity int) bool {
	tmp := make([]int, 1001)
	for _,trip := range(trips){
			tmp[trip[1]] += trip[0]
			tmp[trip[2]] -= trip[0]
	}

	s:=0
	for _,v := range(tmp){
			s += v
			if s > capacity {
					return false
			}
	}
	return true
}