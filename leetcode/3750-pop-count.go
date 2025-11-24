package main
import "math/bits" 
import "fmt" 

func minimumFlips(n int) int {
	minimumFlips := bits.OnesCount((reverse(n) ^ n))
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
