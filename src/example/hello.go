package main

import(
	"fmt"
	"errors"
	"math"
)
type person struct{
	name string
	age int
}
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

	vertices := make(map[string]int)
	vertices["first"] = 1
	vertices["second"] = 2
	fmt.Println(vertices)
	fmt.Println(vertices["first"])
	delete(vertices, "first")
	fmt.Println(vertices)
	fmt.Println(vertices["first"])

	for i:=0; i<5; i++{
		fmt.Println(i)
	}

	i:=0
	for i<5{
		fmt.Println(i)
		i++
	}

	arr4 := []string{"first", "second"}
	for index, value := range arr4{
		fmt.Println("index:",index,"Value:",value)
	} 

	m := make(map[string]int)
	m["first"] = 1
	m["second"] = 2
	for key, value := range arr4{
		fmt.Println("key:",key,"Value:",value)
	} 

	result := sum2(2, 3)
	fmt.Println(result)

	result2, err := sqrt(-16)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result2)
	}

	p := person{name:"sandeep", age:24}
	fmt.Println(p)
	fmt.Println(p.name)

	j := 7
	inc(&j)
	fmt.Println(j)
}

func sqrt(x float64) (float64, error) {
	if x < 0{
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func sum2(x int, y int) int {
	return x + y
}
func inc(x *int) {
	*x++
}