package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	otherNum := 0
	for index, value := range nums {
    	otherNum = target - value
    	if otherNumValue, ok := numsMap[otherNum]; ok {
    		return []int {otherNumValue, index}
		}
		numsMap[value] = index
	}

	return []int {0}
}

func main() {
	nums := []int {3, 3}
	target := 6
	re := twoSum(nums, target)
	fmt.Println(re)
}