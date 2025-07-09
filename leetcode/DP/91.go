package dp

import "strconv"
func numDecodings(s string) int {
    mem := make([]int, len(s)+1)
    for i:=range(mem){
        mem[i] = -1
    }
    var dfs func(int) int 
    dfs = func(i int) (res int){
        if i < 0 {
            return 1
        } 
        p := &mem[i]
        if *p != -1 {
            return *p
        }
        defer func(){*p = res}()
        if s[i] - '0' >= 1{
            res = dfs(i-1)
        }
        if i >= 1{
            str,_ := strconv.Atoi(s[i-1:i+1])
            if str >= 10 && str <= 26{
                res += dfs(i-2)
            }
        } 
        return res

        
    }
    return dfs(len(s)-1)
}