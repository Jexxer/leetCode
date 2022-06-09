package main

import "fmt"

func twoSum(nums []int, target int) []int {
	num_map := make(map[int]int)
	for i, v := range nums {
		if _, ok := num_map[target-v]; !ok {
			num_map[v] = i
		} else {
			return []int{num_map[target-v], i}
		}
	}
	return nil
}

func main() {
	slice := []int{2, 7, 11, 15}
	fmt.Println(twoSum(slice, 18))
}
