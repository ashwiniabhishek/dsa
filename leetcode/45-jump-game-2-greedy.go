package main

import "fmt"


func jump(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
	maxReach := nums[0]
	jumpRequired := 1
	maxReachContenders := nums[0]
    if maxReach >= len(nums)-1 {
        return jumpRequired
    }
	for i:=1;i<len(nums);i++ {
        maxReachContenders = max(i+nums[i], maxReachContenders)
		if maxReach == i {
            maxReach = maxReachContenders
            jumpRequired++
		}
        if maxReach >= len(nums)-1 {
            break
        }
	}
	return jumpRequired
}

func max(a int, b int) int {
	if (a > b) {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
}
