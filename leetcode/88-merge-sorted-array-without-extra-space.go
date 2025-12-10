package main
import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for ;i>=0 && j>=0;k-- {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

	for ;j>=0;k-- {
		nums1[k] = nums2[j]
		j--
	}

	fmt.Println(nums1)
}

func main() {
	merge([]int{1,2,3,0,0,0}, 3, []int{5,6,7}, 3)
	merge([]int{1,2,3,10,0,0,0}, 4, []int{5,6,7}, 3)
	merge([]int{10,20,30,100,0,0,0}, 4, []int{5,6,7}, 3)
	merge([]int{10,20,30,100}, 4, []int{}, 0)
}
