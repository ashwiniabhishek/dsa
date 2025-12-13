package main

import "fmt"

func maxProfit(prices []int) int {
	maxTillNow := prices[len(prices)-1]
	maxProfit := 0
	for i:=len(prices) - 2; i>=0; i-- {
		if prices[i] > maxTillNow {
			maxTillNow = prices[i]
		}
		maxProfit = max(maxProfit, maxTillNow - prices[i])
	}
	return maxProfit
}

func max(a int, b int) int {
	if (a > b) {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
	fmt.Println(maxProfit([]int{4,3,2,1}))
}
