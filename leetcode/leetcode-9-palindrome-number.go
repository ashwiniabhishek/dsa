package main
import "fmt"

func isPalindrome(x int) bool {
	if (x < 0) {
		return false
	} 

	if ( x == 0 ) {
		return true
	}

	return reverse(x) == x
}

func reverse(num int) int {
	rightShiftedNum := num
	currentDigit := 0
	reversedNum := 0
	for rightShiftedNum  > 0 {
		currentDigit = rightShiftedNum % 10
		rightShiftedNum = rightShiftedNum / 10
		reversedNum = reversedNum * 10 + currentDigit 
	}
	fmt.Println("Reversed num: ", reversedNum)
	return reversedNum
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(122))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(1234567887654321))
	fmt.Println(isPalindrome(0))
}
