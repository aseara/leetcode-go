package main

/*
Given an integer array arr, count element x such that x + 1 is also in arr.

If there're duplicates in arr, count them separately.
*/
func countElements(arr []int) (cnt int) {

	var hash [1002]int
	for _, v := range arr {
		hash[v]++
	}

	for _, v := range arr {
		if hash[v+1] > 0 {
			cnt++
		}
	}

	return
}
