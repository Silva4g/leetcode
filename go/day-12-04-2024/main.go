/*
Given n non-negative integers representing an elevation map where the width of
each bar is 1, compute how much water it can trap after raining.

Example 1:
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array
[0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section)
are being trapped.

Example 2:
Input: height = [4,2,0,3,2,5]
Output: 9
*/

package main

import (
	"fmt"
)

func trap(height []int) int {
	if len(height) == 0 {
			return 0
	}

	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	totalWater := 0

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
			leftMax[i] = max(height[i], leftMax[i-1])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
			rightMax[i] = max(height[i], rightMax[i+1])
	}

	for i := 0; i < n; i++ {
			waterAtBar := min(leftMax[i], rightMax[i]) - height[i]
			if waterAtBar > 0 {
					totalWater += waterAtBar
			}
	}

	return totalWater
}

func max(a, b int) int {
	if a > b {
			return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
			return a
	}
	return b
}

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trap([]int{4,2,0,3,2,5}))
	fmt.Println(trap([]int{5,4,3,2,1,0,1,2,3,4,5}))
}