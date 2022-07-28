package main

import (
	"math"
)

func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func twoSum(arr []int, target int) []int {
	for i, v := range arr {
		if j := binarySearch(arr[i+1:], target-v); j != -1 {
			return []int{i, i + 1 + j}
		}
	}
	return []int{-1, -1}
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s1 {
		m[v]++
	}
	for _, v := range s2 {
		m[v]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	odd := 0
	for _, v := range m {
		if v%2 != 0 {
			odd++
		}
	}
	return odd <= 1
}

func isPalindrome2(s string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	odd := 0
	for _, v := range m {
		if v%2 != 0 {
			odd++
		}
	}
	return odd <= 1
}

func leastPerimeter(area int) int {
	sqrt := math.Sqrt(float64(area))
	for i := int(sqrt); i > 0; i-- {
		if area%i == 0 {
			return 2 * (area/i + i)
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(linearSearch([]int{1, 2, 3, 4, 5}, 5))
// 	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 5))
// 	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
// 	fmt.Println(isAnagram("abc", "cba"))
// 	fmt.Println(isPalindrome("abc"))
// 	fmt.Println(isPalindrome2("abc"))
// 	fmt.Println(leastPerimeter([]int{2, 1, 2}))
// }
