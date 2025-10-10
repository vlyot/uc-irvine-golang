package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swap swaps the elements at index i and i+1 in the slice.
func Swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}

// BubbleSort sorts the slice of integers in ascending order using bubble sort.
func BubbleSort(nums []int) {
	n := len(nums)
	for pass := 0; pass < n-1; pass++ {
		for i := 0; i < n-pass-1; i++ {
			if nums[i] > nums[i+1] {
				Swap(nums, i)
			}
		}
	}
}

func main() {
	fmt.Println("Enter up to 10 integers separated by spaces:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fields := strings.Fields(input)

	if len(fields) > 10 {
		fmt.Println("Please enter no more than 10 integers.")
		return
	}

	nums := make([]int, 0, 10)
	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			fmt.Printf("Invalid integer: %s\n", f)
			return
		}
		nums = append(nums, n)
	}

	BubbleSort(nums)

	fmt.Println("Sorted integers:")
	for _, n := range nums {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
