/*
Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

package main

import (
	"fmt"
	"strconv"
)

// Solution
func isPalindrome(x int) bool {
	strNum := strconv.Itoa(x) // Convert to string
	count := 0

	// Slice in ascending order
	sliceAsc := []int{}
	for _, v := range strNum {
		sliceAsc = append(sliceAsc, int(v))
	}

	// Slice in descending order
	sliceDesc := []int{}
	for i := len(sliceAsc) - 1; i >= 0; i-- {
		sliceDesc = append(sliceDesc, sliceAsc[i])
	}

	// Comparing if ascending and descending are equal
	for i := 0; i < len(sliceAsc); i++ {
		if sliceAsc[i] == sliceDesc[i] {
			count++
		}
	}

	return count == len(sliceAsc)
}

func main() {
	// Example
	n := 121
	n2 := 901
	fmt.Println(isPalindrome(n))  // Output true
	fmt.Println(isPalindrome(n2)) // Output false

}
