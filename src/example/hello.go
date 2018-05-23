package main

import(
	"fmt"
)
func main() {
	fmt.Println("Hello, go");	

	var x int = 5
	y := 7
	var sum int = x + y

	fmt.Println(sum)

	if sum < 10{
		fmt.Println("less then 10")
	}else {
		fmt.Println("greater then 10")
	}

	var arr[5] int
	arr[2] = 4
	fmt.Println(arr)

	arr2 := [5]int{1,2,3,4,5}
	fmt.Println(arr2)

	arr3 := []int{1,2,3,4,5}
	arr3 = append(arr3, 6)
	fmt.Println(arr3)
}