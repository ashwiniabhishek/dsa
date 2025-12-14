package main

import "fmt"

func maxProfit(prices []int) int {
    for i := len(prices)-1; i > 0; i-- {
        prices[i] = prices[i] -  prices[i-1]
    }
    
    prices[0] = 0
    profit := prices[0]
    
    fmt.Println(prices)
    
    for i := 1; i < len(prices); i++ {
        if prices[i] < prices[i] + prices[i-1] {
            prices[i] = prices[i] + prices[i-1]
        }
        
        if profit < prices[i] {
                profit = prices[i]
        }
    }
    
    return profit
}

func main() {

}
