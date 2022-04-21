package main

import (
	"fmt"
	"sort"
)

// func printSliceInfo(slc []int) {
// 	fmt.Println("len: ", len(slc), " cap: ", cap(slc))
// 	fmt.Println("Content: ", slc)
// }

// func printArrayInfo(arr [3]int) {
// 	fmt.Println("len: ", len(arr), " cap: ", cap(arr))
// 	fmt.Println("Content: ", arr)
// }

// func printMap(m map[string]int) {
// 	fmt.Println(m)
// 	fmt.Println("len: ", len(m))
// }
// type human struct {
// 	name string
// 	age int
// }
// type student struct {
// 	human
// 	grade string
// }

type int_set map[int]bool

func printIntSet(set int_set) {
	sli := []int{}
	for k := range set{
		sli = append(sli, k)
	}
	sort.Ints(sli)
	fmt.Println(sli)
}

func initIntSet(set int_set, sli []int) {
	for _, v := range sli{
		set[v] = true
	}
}

func addIntSet(set int_set, i int) bool {
	set[i] = true
	return set[i]
}

func removeIntSet(set int_set, i int) {
	delete(set, i)
}

func isInIntSet(set int_set, i int) bool {
	_, ok := set[i]
	return ok
}

func main() {
	// arr3 := [3]int{1,2,3}
	// arr3_2 := [3][2]int{{1,2}, {3,4},{5,1}}
	// var arr3_4 [3][4]int
	// printArrayInfo(arr3)
	// fmt.Println((arr3[0]))
	// fmt.Println(cap(arr3_2[0]))
	// fmt.Println(arr3_4)

	// var slc3 = []int{1, 2, 3}
	// fmt.Println(slc3)
	// fmt.Println(len(slc3), cap(slc3))

	// var slc_nil []int
	// fmt.Println(slc_nil == nil)
	// fmt.Printf("%T\n", arr3)
	// fmt.Printf("%T\n", slc_nil)
	// printSliceInfo(slc_nil)
	// slc_nil = append(slc_nil, 1)
	// printSliceInfo((slc_nil))
	// slc_nil = append(slc_nil, 2, 3)
	// printSliceInfo((slc_nil))
	// slc_nil = append(slc_nil, 4, 5)
	// printSliceInfo((slc_nil))
	// new_slc := make([]int, 0, 5)
	// printSliceInfo(new_slc)
	// new_slc = append(new_slc, 2)
	// printSliceInfo(new_slc)
	// var size_slc = 3
	// new_slc1 := make([]int, size_slc)
	// printSliceInfo((new_slc1))
	// var non_nil_slc = []int{}
	// fmt.Println(non_nil_slc == nil)

	// You could think that an array is the special case of slice which has fixed len and cap (they are equal). and the slice is the "dynamic" array (len is less or euqal to cap)
	// Remember that array and slice in go, both of their type include their length and capacity
	// make method only make slice but not arry (arry has fixed length and capacity)

	// slc_orig := []int{1,2,3,4,5,6}
	// printSliceInfo(slc_orig)
	// slc_cut1 := slc_orig[:]
	// printSliceInfo(slc_cut1)
	// slc_cut2 := slc_orig[2:]
	// printSliceInfo(slc_cut2)
	// the third parameter z of slice[x:y:z] will create a new slice in memory instead of just copying the pointer from the original slice (whic will change both if changed one of them)
	// slc_cut3 := slc_orig[:3:3]
	// printSliceInfo(slc_orig)
	// printSliceInfo(slc_cut3)

	// slc_cut3 = append(slc_cut3, 999)
	// printSliceInfo(slc_orig)
	// printSliceInfo(slc_cut3)

	// m1 := map[string]int{}
	// printMap(m1)
	// m1["hello"] = 2
	// printMap(m1)
	// m2 := map[string]int{
	// 	"Hello" : 1,
	// 	"World" : 2,
	// }
	// for k, _ := range m2 {
	// 	fmt.Print(k, " ")
	// }
	// fmt.Printf("\b!\n")
	// delete(m2, "Hello")
	// printMap((m2))

	// hemos := human{
	// 	"Jason",
	// 	10,
	// }
  // kid := student{
	// }
	// fmt.Println(kid)
	// kid.age = hemos.age
	// kid.name = hemos.name
	// kid.grade = "K9"
	// fmt.Println(kid.name)
	iset := int_set{}
	printIntSet(iset)
	sli := []int{1,2,3,4,5}
	initIntSet(iset, sli)
	printIntSet(iset)
	addIntSet(iset, 5)
	printIntSet(iset)
	fmt.Println(isInIntSet(iset, 4))
	fmt.Println(isInIntSet(iset, 8))
	addIntSet(iset, 8)
	printIntSet(iset)
	removeIntSet(iset, 8)
	printIntSet(iset)
}