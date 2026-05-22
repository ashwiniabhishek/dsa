package main

import "fmt"

func main() {

}

func productExceptSelf(nums []int) []int {

	n := len(nums)
	prefixProduct := make(n+1, []int)
	prefixProduct[0] = 1 
	suffixProduct := make(n+1, []int)
	suffixProduct[n] = 1
	temp := 1
	for i:=1;i<=n;i++ {
		temp = temp*nums[i-1]
		prefixProduct[i] = temp 
	}

	temp = 1

	for i:=n-2;i>=0;i-- {
		temp = temp*nums[i+1]
		suffixProduct[i] = temp
	}

	answer := make(n, []int)

	for i:=0;i<n;i++ {
		answer[i] = prefixProduct[i]*suffixProduct[i]
	}
}
