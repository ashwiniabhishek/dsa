package main
import "fmt"


func plusOne(digits []int) []int {
	borrow := 0
	for i:= len(digits)-1; i>=0; i-- {
		currentDigit := 0
		sum := 0
		if i == len(digits)-1 {
			sum = digits[i] + 1	
		} else {
			sum = digits[i] + borrow
		}
		if (sum > 9) {
			borrow = 1
			currentDigit = 0 
		} else {
			borrow = 0
			currentDigit = sum
		}
		digits[i] = currentDigit
	}

	arr := [1]int{borrow}

	if borrow == 0 {
		return digits
	} else {
		return append(arr[:], digits...)
	}
}

func main() {
	fmt.Println(plusOne([]int{1,2,3}))
	fmt.Println(plusOne([]int{9,9,9}))
}
