/***************************************
Getting Started with Golang, Hands On

--tvbgeek
****************************************/

package main

import (
	"fmt"
)

// maintains the order of elements
// O(n)
func removeOrdered(a []int, i int) ([]int) {
	if len(a) == 0 {
		return a
	}
	return append(a[:i], a[i+1:]...)
}

// does not maintain the order of the elements
// O(1)
func removeUnordered(a []int, i int) ([]int) {
	if len(a) == 0 {
		return a
	}
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func main() {

	// create a slice
	arr := make([]int, 2)
	fmt.Println("slice", arr)

	// create a slice of size 0 and try accessing it 
	tmp := make([]int, 0)
	// the following assignments will throw error, index out of bounds
	// tmp[0] = 1
	// tmp[1] = 2
	// though you can use append 
	tmp = append(tmp, 1)
	tmp = append(tmp, 2)
	
	// update value at any index
	arr[0] = 1
	fmt.Println("slice after updating value at 0th index", arr)

	// read value at any index
	fmt.Println("element at 0th index", arr[0])

	// append a value
	arr = append(arr, 3)
	fmt.Println("slice after append", arr)

	// number of elements in a slice
	fmt.Println("total elements in slice", len(arr))

	// creating a new slice and copying
	arr_1 := make([]int, len(arr))
	copy(arr_1, arr)
	fmt.Println("newly copied slice", arr_1)

	// appending a slice to slice
	arr = append(arr, arr_1...)
	fmt.Println("appending a slice", arr)

	// providing intial elements of the slice
	arr = []int{0, 1, 2, 3, 4}
	fmt.Println("our new slice", arr)

	// iterating over slice elements
	for i, val := range arr {
		fmt.Println("index:", i, "val:", val)
	}

	//Slicing the slice :)
	/*
		https://stackoverflow.com/questions/509211/understanding-slice-notation
		a[start:stop]  # items start through stop-1
		a[start:]      # items start through the rest of the array
		a[:stop]       # items from the beginning through stop-1
		a[:]           # a copy of the whole array
	*/

	fmt.Println("complete slice", arr[:])
	fmt.Println("last element", arr[len(arr)-1])
	fmt.Println("elements from second index (second index is included)", arr[2:])
	fmt.Println("elements till second index (second index is not included)", arr[:2])
	fmt.Println("elements 1st to 4th index (4th index is not included)", arr[1:4])

	//complete list
	fmt.Println("complete slice", arr)

	//remove an element
	//https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang/37335777#37335777
	// O(n) remove maintains the ordering of the elements
	arr = removeOrdered(arr, 0)
	fmt.Println("remove element at 0th index ordered", arr)

	// O(1) remove does not maintain the ordering of the elements
	arr = removeUnordered(arr, 0)
	fmt.Println("remove element at 0th index unordered", arr)

	// deletign a slice 
	arr = nil
	fmt.Println("delete slice completely", arr)
}
