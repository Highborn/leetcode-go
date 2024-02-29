package problems

import (
	"log"
	"sort"
)

func ArrayRankTransform(arr []int) []int {
	refArr := make([]*int, len(arr))

	for i := range arr {
		refArr[i] = &(arr[i])
	}
	// QuickSort(&refArr, 0, len(arr)-1)
	sort.Slice(refArr, func(i, j int) bool {
		return *refArr[i] < *refArr[j]
	})
	lastSeen := int(1e10)
	rnk := 0
	for _, ptr := range refArr {
		log.Print(*ptr)
		if *ptr == lastSeen {
			*ptr = rnk
		} else {
			rnk += 1
			lastSeen = *ptr
			*ptr = rnk
		}
	}

	return arr
}

func QuickSort(pArr *[]*int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(pArr, start, end)
	QuickSort(pArr, start, pivot-1)
	QuickSort(pArr, pivot+1, end)
}

func partition(pArr *[]*int, start, end int) int {
	pivot := (*pArr)[end]
	idx := start
	for i, pointer := range (*pArr)[start:end] {
		if *pointer < *pivot {
			(*pArr)[i+start], (*pArr)[idx] = (*pArr)[idx], (*pArr)[i+start]
			idx += 1
		}
	}
	// idx += 1
	(*pArr)[end], (*pArr)[idx] = (*pArr)[idx], (*pArr)[end]
	return idx
}

func ArrayRankTransform2(arr []int) []int {
	arrClone := make([]int, len(arr))
	copy(arrClone, arr)
	rnkMap := make(map[int]int)
	sort.Ints(arrClone)
	rnk := 1
	for _, item := range arrClone {
		if rnkMap[item] == 0 {
			rnkMap[item] = rnk
			rnk += 1
		}
	}
	for i, item := range arr {
		arrClone[i] = rnkMap[item]
	}
	return arrClone
}
