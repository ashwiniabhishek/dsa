package main

import "fmt"

func canJump(nums []int) bool {
	for i:=0;i<len(nums)-1;i++ {
		if nums[i] == 0 {
			jumpPossible := false
			for j:=i-1;j>=0;j-- {
				jumpCompensation := i-j
				if nums[j] >= 1 + jumpCompensation {
					jumpPossible = true
					break
				}
			}
			if !jumpPossible {
				return jumpPossible
			}
		}
	}
	return true
}
func main() {
	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))
	fmt.Println(canJump([]int{3}))
	fmt.Println(canJump([]int{5,0,0,0,0,1,0}))
	fmt.Println(canJump([]int{4,1,1,0,0,0,0}))
}
