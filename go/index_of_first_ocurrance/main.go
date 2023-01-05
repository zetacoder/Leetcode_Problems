/*
28. Find the Index of the First Occurrence in a String
Medium

// Instructions
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack,
or -1 if needle is not part of haystack.


Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.


Constraints:

1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.
*/

package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	hLen := len(haystack)
	nLen := len(needle)
	for i := 0; i < hLen-nLen+1; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}
	return -1
}

func main() {
	// Test 1
	haystack1 := "hello"
	needle1 := "ll"

	// Test 2
	haystack2 := "sadbutsad"
	needle2 := "sad"

	fmt.Println(strStr(haystack1, needle1)) // Output: 2 --> There is needle into haystack, first ocurrence in index 2 "l"
	fmt.Println(strStr(haystack2, needle2)) // Output: 0 --> There is needle into haystack, first ocurrence in index 0 "s"

}

/*
Success
Details
Runtime: 1 ms, faster than 54.86% of Go online submissions for Find the Index of the First Occurrence in a String.
Memory Usage: 1.9 MB, less than 88.40% of Go online submissions for Find the Index of the First Occurrence in a String.
*/
