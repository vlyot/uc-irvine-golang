package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// partitionArray divides the array into n approximately equal parts.
func partitionArray(arr []int, n int) [][]int {
	length := len(arr)
	partitions := make([][]int, 0, n)
	chunkSize := length / n
	remainder := length % n
	start := 0

	for i := 0; i < n; i++ {
		end := start + chunkSize
		if remainder > 0 {
			end++
			remainder--
		}
		partitions = append(partitions, arr[start:end])
		start = end
	}
	return partitions
}

// mergeSortedArrays merges multiple sorted arrays into one sorted array.
func mergeSortedArrays(arrays [][]int) []int {
	result := make([]int, 0)
	indices := make([]int, len(arrays))

	for {
		minIndex := -1
		minValue := 0
		for i, arr := range arrays {
			if indices[i] < len(arr) {
				if minIndex == -1 || arr[indices[i]] < minValue {
					minIndex = i
					minValue = arr[indices[i]]
				}
			}
		}
		if minIndex == -1 {
			break
		}
		result = append(result, minValue)
		indices[minIndex]++
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a series of integers separated by spaces:")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strNums := strings.Fields(input)

	if len(strNums) == 0 {
		fmt.Println("No integers provided.")
		return
	}

	nums := make([]int, len(strNums))
	for i, s := range strNums {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Invalid integer: %s\n", s)
			return
		}
		nums[i] = n
	}

	const partitionCount = 4
	partitions := partitionArray(nums, partitionCount)
	sortedPartitions := make([][]int, partitionCount)
	done := make(chan int)

	for i := 0; i < partitionCount; i++ {
		go func(idx int, subarr []int) {
			fmt.Printf("Goroutine %d sorting subarray: %v\n", idx+1, subarr)
			sort.Ints(subarr)
			sortedPartitions[idx] = subarr
			done <- idx
		}(i, partitions[i])
	}

	// Wait for all goroutines to finish
	for i := 0; i < partitionCount; i++ {
		<-done
	}

	merged := mergeSortedArrays(sortedPartitions)
	fmt.Printf("Sorted array: %v\n", merged)
}
