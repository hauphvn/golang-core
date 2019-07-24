package main

import "fmt"

func main() {
	// Spice
	var arrInt = [5]int{22, 33, 44, 55, 66}
	fmt.Println(arrInt)
	var sliceInt []int
	fmt.Println(sliceInt)
	sliceInt = arrInt[:]
	fmt.Println(sliceInt)
	sliceInt = arrInt[:2]
	fmt.Println(sliceInt)

	fmt.Println("length vs capacity")
	sliceInt = arrInt[1:4]
	sliceInt[0] = 3333
	fmt.Println(arrInt)
	fmt.Println(cap(sliceInt))
	fmt.Println(len(sliceInt))
	//sliceInt[3] = 999 overflow

	fmt.Println("Methods of Slice")
	var sliceInt2 = make([]int, 2, 5)
	fmt.Println(len(sliceInt2))
	fmt.Println(cap(sliceInt2))
	fmt.Println(sliceInt2)

	sliceInt2 = append(sliceInt2, 999)
	sliceInt2 = append(sliceInt2, 999)
	sliceInt2 = append(sliceInt2, 999)
	sliceInt2 = append(sliceInt2, 999)

	sliceInt2 = append(sliceInt2, 999)
	sliceInt2 = append(sliceInt2, 999)
	sliceInt2 = append(sliceInt2, 999)
	sliceInt2 = append(sliceInt2, 999)
	sliceInt2 = append(sliceInt2, 999)
	fmt.Println(cap(sliceInt2))
	fmt.Println(sliceInt2)

	var arrInt3 = []int{1, 2, 3}
	var sliceInt4 = make([]int, 3)
	fmt.Println(cap(sliceInt4))
	num := copy(sliceInt4, arrInt3)
	fmt.Println(cap(sliceInt4))
	fmt.Println(num)
	fmt.Println(sliceInt4)

	// Delete element in slice
	var sliceInt5 = []int{1, 2, 3, 4, 5}
	var sliceInt6 = append(sliceInt5[0:1], sliceInt5[2:]...)
	fmt.Println(sliceInt6)

}
