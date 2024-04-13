/*
Given a rows x cols binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.

Example 1:
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 6
Explanation: The maximal rectangle is shown in the above picture.

Example 2:
Input: matrix = [["0"]]
Output: 0

Example 3:
Input: matrix = [["1"]]
Output: 1
*/

package main

import (
	"fmt"
)

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	maxArea := 0
	cols := len(matrix[0])
	height := make([]int, cols)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				height[j] += 1
			} else {
				height[j] = 0
			}
		}
		maxArea = max(maxArea, largestRectangleArea(height))
	}
	return maxArea
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := []int{}
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, h*w)
		}
		stack = append(stack, i)
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
			return x
	}
	return y
}


func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
}
fmt.Println("Maximum rectangle area:", maximalRectangle(matrix))
}