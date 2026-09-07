class Solution {
    public int singleNumber(int[] nums) {
        int answer = 0;
        for(int num : nums) {
            answer = answer ^ num;
        }
        return answer;
    }
}


/***
 
Why XOR works ?

XOR (^) has three important properties:

a ^ a = 0
A number XORed with itself becomes 0.

a ^ 0 = a
XORing with 0 leaves the number unchanged.

XOR is commutative and associative, thus we can xor in any permutation.
a ^ b = b ^ a
(a ^ b) ^ c = a ^ (b ^ c)

***/

