package main

import "fmt"

func hIndex(citations []int) int {
    	h := 0
	for i:=0;i<len(citations);i++ {
		count := i+1
		for j:=0;j<len(citations);j++ {
			if (citations[j] >= i+1) {
				count--
			}
		}
		if count <= 0 {
			h = max(h, i+1)
		}

	}
	return h
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

	fmt.Println(hIndex([]int{3,0,6,1,5}))
	fmt.Println(hIndex([]int{1,3,1}))
}
