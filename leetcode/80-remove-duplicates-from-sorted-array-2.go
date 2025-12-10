package main
import "fmt"

func removeDuplicates(nums []int) int {
	currentValue := nums[0]
	frequency := 0
	slowPointer := 0
	
	for fastPointer := 0;fastPointer < len(nums);fastPointer++ {
		if currentValue == nums[fastPointer] {
			frequency++
		} else {
			currentValue = nums[fastPointer]
			frequency = 1
		}

		if frequency <= 2 {
			nums[slowPointer] = nums[fastPointer]
			slowPointer++
		}
	}
	fmt.Println(nums, slowPointer)
	return slowPointer
}

func main() {
	removeDuplicates([]int{1,1,1,2,2,3,3,3,4})
}
