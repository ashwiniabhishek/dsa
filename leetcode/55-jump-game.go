package main

import "fmt"

func canJump(nums []int) bool {
	dp := make([]bool, len(nums)) 
	dp[len(nums)-1] = true
	for i:=len(nums)-2;i>=0;i-- {
		for j:=i+1; j<=i+nums[i] && j<len(nums); j++ {
			dp[i] = dp[j] || dp[i]
		}
	}
	return dp[0]
}


func main() {
	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))
	fmt.Println(canJump([]int{3}))
	fmt.Println(canJump([]int{5,0,0,0,0,1,0}))
}
