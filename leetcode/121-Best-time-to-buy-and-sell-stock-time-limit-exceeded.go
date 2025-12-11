package main

import "fmt"

func maxProfit(prices []int) int {
    maxProfit := 0
    for i:=0;i<len(prices);i++ {
	    for j:=i+1;j<len(prices);j++ {
		    maxProfit = max(maxProfit, prices[j]-prices[i])
		    // fmt.Println(maxProfit)
	    }
    }
    return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}

func max(a int, b int) int {
	if a  > b {
		return a
	} else {
		return b
	}
}
