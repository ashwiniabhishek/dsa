package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
	fmt.Println(nums)
}

func reverse(nums []int) {
	for i,j:= 0,len(nums)-1; i<j; i,j = i+1,j-1 {
		nums[i],nums[j] = nums[j],nums[i]
	}
}

func main() {
	rotate([]int{1,2,3,4,5}, 3)
}
