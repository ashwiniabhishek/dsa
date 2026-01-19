package main

import "fmt"

func hIndex(citations []int ) int {
	arr := make([]int, len(citations)+1)
	for i:=0;i<len(citations);i++ {
		if citations[i] <= len(citations) {
			arr[citations[i]] +=  1 
		} else {
			arr[len(citations)] += 1
		}
	}

	count := 0
	fmt.Println(arr)

	for i:=len(citations);i>=1;i-- {
		count += arr[i]
		if count >= i {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(hIndex([]int{3,0,6,1,5}))
	fmt.Println(hIndex([]int{1,3,1}))
}
