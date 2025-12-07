package main

import "fmt"

func rotate(nums []int, k int) {
	numsSize := len(nums)
	fmt.Println("numsSize: ", numsSize)
	rotatedArray := make([]int, numsSize)
	for i:=0;i<numsSize;i++ {
		rotatedArrayIndex := (i + k) % numsSize
		rotatedArray[rotatedArrayIndex] = nums[i]
	}
	copy(nums,rotatedArray)
	fmt.Println(nums)

}

func main() {
	rotate([]int{1,2,3}, 2)
	rotate([]int{1,2,3}, 3)
	rotate([]int{1,2,3,4,5,6,7}, 3)
}
