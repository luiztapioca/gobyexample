package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func redefineMap(old *[]int, new ...int) {
	*old = new
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("old:", testSlice)
	redefineMap(&testSlice, []int{2, 3, 4, 5}...)
	fmt.Println("new:", testSlice)
}
