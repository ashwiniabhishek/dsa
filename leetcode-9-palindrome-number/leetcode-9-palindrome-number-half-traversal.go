package main
import "fmt"

func isPalindrome(x int) bool {
	// not able to fully grasp the 2nd condition after or
	if (x < 0 || (x % 10 == 0 && x != 0)) {
		return false
	} 

	return isPalindromeUsingHalfTraversal(x)
}

func isPalindromeUsingHalfTraversal(num int) bool {
	rightShiftedNum := num
	currentDigit := 0
	reversedNum := 0
	for rightShiftedNum  > reversedNum {
		currentDigit = rightShiftedNum % 10
		rightShiftedNum = rightShiftedNum / 10
		reversedNum = reversedNum * 10 + currentDigit 
	}
	fmt.Println("Reversed num: ", reversedNum)
	return rightShiftedNum == reversedNum || rightShiftedNum == (reversedNum/10)
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(122))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(1234567887654321))
	fmt.Println(isPalindrome(10))
}
