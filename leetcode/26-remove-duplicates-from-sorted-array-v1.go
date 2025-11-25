package main

import "fmt"

func removeDuplicates(nums []int) int {
    currentUniqueNumIndex := 0
    
    for i:=1;i<len(nums);i++ {
	    if nums[currentUniqueNumIndex] != nums[i] {
		    currentUniqueNumIndex++
		    nums[currentUniqueNumIndex] = nums[i]
	    }
    }
    fmt.Println(nums)
    return currentUniqueNumIndex + 1
}


func main() {
	fmt.Println(removeDuplicates([]int{1,1,1,2,3,3,3,4,5}))
	fmt.Println(removeDuplicates([]int{1,1,1,2,2,2,2,2,2,2,2,2,2}))
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{}))
}
