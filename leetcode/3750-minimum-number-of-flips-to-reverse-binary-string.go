package main
import "fmt" 

func minimumFlips(n int) int {
	reversedBinary := reverse(n)
	xor := reversedBinary ^ n
	minimumFlips := 0
	bit := 0
	for xor > 0 {
		bit = xor & 1
		minimumFlips += bit
		xor = xor >> 1
	}
	fmt.Println("flips: ", minimumFlips)
	
	return minimumFlips 
}

func reverse(n int) int {
	bit := 0
	rightShiftedBinary := n
	reversedBinary := 0

	for rightShiftedBinary > 0 {
		bit = rightShiftedBinary & 1
		rightShiftedBinary = rightShiftedBinary >> 1
		reversedBinary = (reversedBinary << 1) | bit

	}
	fmt.Println("reversed number: ", reversedBinary)
	return reversedBinary
}

func main() {
	minimumFlips(7)
	minimumFlips(10)
}
