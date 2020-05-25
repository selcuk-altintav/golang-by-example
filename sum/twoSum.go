// https://leetcode.com/problems/two-sum/

package main

import (
	"fmt"
)

func main() {
	var nums [5]int = [5]int{2, 91, 11, 7, 12}
	target := 9
	fmt.Println(twoSum(nums[:], target))
}

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		firstHalf := target - nums[i]
		if _, ok := indexMap[firstHalf]; ok {
			return []int{indexMap[firstHalf], i}
		}
		indexMap[nums[i]] = i
	}
	return nil
}
