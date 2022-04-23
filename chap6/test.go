package main

import (
	"fmt"
)

type simpleStruct struct {
	name string
} 

func testFunc(ad1 string, ad2 string, typ string) {
	if ad1 != ad2 {
		fmt.Printf("%s is passed by value of copy\n", typ)
	} else {
		fmt.Printf("%s is passed by the value of the pointer\n", typ)
	}
}
// test functions for basic type

func intFunc(i int, ia *int){
		ad1 := fmt.Sprintf("%p", &i)
		ad2 := fmt.Sprintf("%p", ia)
		testFunc(ad1, ad2, "int")
}

func floatFunc(f float64, fa *float64){
		ad1 := fmt.Sprintf("%p", &f)
		ad2 := fmt.Sprintf("%p", fa)
		testFunc(ad1, ad2, "float64")
}

func boolFunc(b bool, ba *bool){
		ad1 := fmt.Sprintf("%p", &b)
		ad2 := fmt.Sprintf("%p", ba)
		testFunc(ad1, ad2, "boolean")
}

func stringFunc(s string, sa *string){
		ad1 := fmt.Sprintf("%p", &s)
		ad2 := fmt.Sprintf("%p", sa)
		testFunc(ad1, ad2, "string")
}

func complexFunc(c complex128, ca *complex128){
		ad1 := fmt.Sprintf("%p", &c)
		ad2 := fmt.Sprintf("%p", ca)
		testFunc(ad1, ad2, "complex128")}

// test function for the Array (int) with length 5 (fixed)

func arrayFunc(a [5]int, ad0 *int) [5]int{
	ad1 := fmt.Sprintf("%p", &a[0])
	ad2 := fmt.Sprintf("%p", ad0)
	testFunc(ad1, ad2, "Array");
	a[2] = 666
	return a
}

// test function for the Slice (int)
func sliceFunc(s []int, ad0 *int) []int{
	ad1 := fmt.Sprintf("%p", s)
	ad2 := fmt.Sprintf("%p", ad0)
	testFunc(ad1, ad2, "Slice")
	s[2] = 666
	s = append(s, 6)
	return s
}

// test function for the map [string]int

func mapFunc(m map[string]int, p *map[string]int){
	ad1 := fmt.Sprintf("%p", m)
	ad2 := fmt.Sprintf("%p", *p)
	testFunc(ad1, ad2, "Map")
	m["Hello"] = 666
	m["!"] = 999
}

// test function for struct

func structFunc(s simpleStruct, p *simpleStruct) {
	ad1 := fmt.Sprintf("%p", &s)
	ad2 := fmt.Sprintf("%p", p)
	testFunc(ad1, ad2, "Struct")
	fmt.Println(s)
	s.name = "WOrld"
	fmt.Println(s)
}

func main() {
	var i int = 2
	intFunc(i, &i)
	var f float64 = 1.2345
	floatFunc(f, &f)
	var b bool = true
	boolFunc(b, &b)
	var str string = "Hello!"
	stringFunc(str, &str)
	var comp complex128 = 123 + 2i
	complexFunc(comp, &comp)

	//array 
	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr)
	fmt.Printf("%p\n", &arr) // you must referrenced the array variable since array name is NOT the pointer as in other language (like C, C++, JAVA)
	arr_r := arrayFunc(arr, &arr[0])
	fmt.Println(arr_r)
	fmt.Println(arr)

	//slice
	slc := []int{1,2,3,4,5}
	fmt.Printf("%p\n", slc) // In slice, the variable name is the pointer to the first element
	fmt.Printf("%p\n", &slc[0]) // this should be the same as above
	fmt.Println(slc)
	slc_r :=sliceFunc(slc, &slc[0])

	// content changed but the len and cap is not change, so the append operation in the function is not reflect in the original 
	fmt.Println(slc)
	fmt.Println(slc_r)

	m := map[string]int{
		"Hello": 1,
		"World": 2,
	}

	fmt.Printf("%p\n", m);
	fmt.Println(m["Hello"])
	fmt.Println(m["!"])
	fmt.Println(len(m))
	mapFunc(m, &m) // this is not wrong, since the map itself is the pointer, then &m is the pointer to the pointer, which could passed the address
	fmt.Println(m["Hello"])
	fmt.Printf("%p\n", m);
	fmt.Println(m["!"])
	fmt.Println(len(m)) // len could change

	stru := simpleStruct{"Hello"}
	structFunc(stru, &stru)
	// not change, passed by value of copy 
	fmt.Println(stru)
}