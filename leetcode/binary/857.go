package main

import "slices"
func minEatingSpeed(piles []int, h int) int {
    n := len(piles)
    left, right := 0, slices.Max(piles)
    condiInt := sumN(piles)
    for left + 1 < right{
        mid := (left+right)/2
        if condiInt/mid <= h -n {
            right = mid
        }else{
            left = mid
        }
    }
    return right

}

func sumN(piles []int) int{
    sum := 0
    for i:=range(piles){
     sum += i   
    }
    return sum -len(piles)
}


func main(){
	minEatingSpeed([]int{3,6,7,11}, 8)
}