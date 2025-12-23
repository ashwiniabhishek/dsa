package main

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[len(prices)-1][0] = 0
	dp[len(prices)-1][1] = prices[len(prices)-1]
	for i:=len(prices)-2;i>=0;i-- {
		dp[i][0] = max(-prices[i] + dp[i+1][1], dp[i+1][0]) 
		dp[i][1] = max(prices[i] + dp[i+1][0], dp[i+1][1]) 
	}
	return dp[0][0]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}
