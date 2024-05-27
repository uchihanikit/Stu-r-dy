package main

import (
	"fmt"
	"math"
)

type Dictionary map[int]int

func (d Dictionary) Search(num int) bool {
	_, ok := d[num]
	if ok {
		return true
	}
	return false
}

func containsDuplicate(nums []int) bool {
	dictionary := Dictionary{}
	for _, num := range nums {
		if dictionary.Search(num) {
			return true
		}
		dictionary[num]++
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	dictionary := make(map[int]int)
	for _, num := range nums {
		_, ok := dictionary[num]
		if ok {
			return true
		}
		dictionary[num]++
	}
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	dictionary := Dictionary{}
	for i := 0; i < len(nums); i++ {
		//fmt.Println(i, nums[i], dictionary[nums[i]])
		if dictionary.Search(nums[i]) {
			//fmt.Println(" ", int(math.Abs(float64(dictionary[nums[i]]-i))), k)
			if int(math.Abs(float64(dictionary[nums[i]]-i))) <= k {
				return true
			}
		}
		dictionary[nums[i]] = i
	}
	return false
}

//func containsNearbyDuplicate(nums []int, k int) bool {
//	for i := 0; i <= len(nums); i++ {
//		for j := i + 1; j <= len(nums)-1; j++ {
//			fmt.Println(nums[i], nums[j], nums[i] == nums[j] && int(math.Abs(float64(i-j))) <= k)
//			if nums[i] == nums[j] && int(math.Abs(float64(i-j))) <= k {
//				return true
//			}
//		}
//	}
//	return false
//}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
