package main

import "fmt"

func removeDuplicates(nums []int) int {
    slowPointer := 0
    
    for fastPointer:=1;fastPointer < len(nums);fastPointer++ {
	    if nums[slowPointer] != nums[fastPointer] {
		    slowPointer++
		    nums[slowPointer] = nums[fastPointer]
	    }
    }
    fmt.Println(nums)
    return slowPointer + 1
}


func main() {
	fmt.Println(removeDuplicates([]int{1,1,1,2,3,3,3,4,5}))
	fmt.Println(removeDuplicates([]int{1,1,1,2,2,2,2,2,2,2,2,2,2}))
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{}))
}
