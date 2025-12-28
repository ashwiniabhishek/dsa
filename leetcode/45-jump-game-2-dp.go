package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	dp := make([]int, len(nums)) 
	dp[len(nums)-1] = 0 
	for i:=len(nums)-2;i>=0;i-- {
		dp[i] = math.MaxInt
		for j:=i+1; j<=i+nums[i] && j<len(nums); j++ {
			dp[i] = min(addOne(dp[j]),dp[i])
		}
	}
	return dp[0]
}

func min(a int, b int) int {
	if (a < b) {
		return a
	} else {
		return b
	}
}

func addOne(num int) int {
	if num == math.MaxInt {
		return num
	} else {
		return 1 + num
	}
}


func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
	fmt.Println(jump([]int{2,3,0,1,4}))
	fmt.Println(jump([]int{3}))
	fmt.Println(jump([]int{5,0,0,0,0,1,0}))
}
