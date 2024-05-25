package main

import "fmt"

func main() {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	output := removeDuplicatesSolution1(input)
	fmt.Println(output)
}

func removeDuplicatesSolution1(nums []int) int {
	mapper := make(map[int]int)

	i := 0
	for _, num := range nums {
		if _, ok := mapper[num]; ok {
			continue
		} else {
			mapper[num] = num
			nums[i] = num
			i++
		}
	}

	return len(mapper)
}
