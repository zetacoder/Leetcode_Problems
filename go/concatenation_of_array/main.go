/*
1929. Concatenation of Array

Given an integer array nums of length n,
you want to create an array ans of length 2n where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).

Specifically, ans is the concatenation of two nums arrays.

Return the array ans.

Example 1:

Input: nums = [1,2,1]
Output: [1,2,1,1,2,1]
Explanation: The array ans is formed as follows:
- ans = [nums[0],nums[1],nums[2],nums[0],nums[1],nums[2]]
- ans = [1,2,1,1,2,1]
Example 2:

Input: nums = [1,3,2,1]
Output: [1,3,2,1,1,3,2,1]
Explanation: The array ans is formed as follows:
- ans = [nums[0],nums[1],nums[2],nums[3],nums[0],nums[1],nums[2],nums[3]]
- ans = [1,3,2,1,1,3,2,1]


Constraints:

n == nums.length
1 <= n <= 1000
1 <= nums[i] <= 1000
*/

package main

import "fmt"

// SOLUTION
func getConcatenation(nums []int) []int {
	ans := nums

	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[i]) // Append nums to ans.
	}
	return ans
}

func main() {

	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{3, 2, 1, 0}
	nums3 := []int{0, 1, 0, 1, 0}

	fmt.Println(getConcatenation(nums1)) // Expected: [1 2 3 4 5 1 2 3 4 5] / Output: [1 2 3 4 5 1 2 3 4 5]
	fmt.Println(getConcatenation(nums2)) // Expected: [3 2 1 0 3 2 1 0] / Output: [3 2 1 0 3 2 1 0]
	fmt.Println(getConcatenation(nums3)) // Expected: [0 1 0 1 0 0 1 0 1 0] / Output: [0 1 0 1 0 0 1 0 1 0]

}

/*
Success
Details 
Runtime: 33 ms, faster than 8.07% of Go online submissions for Concatenation of Array.
Memory Usage: 6.3 MB, less than 62.42% of Go online submissions for Concatenation of Array.
*/
