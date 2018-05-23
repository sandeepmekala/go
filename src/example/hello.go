package main

import(
	"fmt"
	"errors"
	"math"
	"strings"
	"sort"
)
type person struct{
	name string
	age int
}
type Rectange struct{
	leftX float64
	topY float64
	width float64
	height float64
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
	arr3 = arr3[2:6]
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

	const PI float64 = 3.1435343
	fmt.Println(PI)

	yourAge := 12
	switch yourAge{
		case 10: fmt.Println("you are 10")
		case 12: fmt.Println("you are 12")
		case 16: fmt.Println("you are 16")
		default: fmt.Println("you are 16")
	}

	fmt.Println(safeDiv(3,0))
	fmt.Println(safeDiv(3,2))

	yPtr := new(int)
	changeYNow(yPtr)
	fmt.Println(*yPtr)

	rect := Rectange{10,20,30,40}
	fmt.Println("area: ",rect.area())

	something := "Hello world"
	fmt.Println(strings.Contains(something, "l"))
	fmt.Println(strings.Index(something, "l"))
	fmt.Println(strings.Count(something, "l"))
	fmt.Println(strings.Replace(something, "l", "x",3))

	csvString := "1,2,3,4,5"
	fmt.Println(strings.Split(csvString, ","))

	listOfLetters := []string {"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println(listOfLetters)

	listOfNums := strings.Join([]string{"1","2","3"}, ",")
	fmt.Println(listOfNums)

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
func safeDiv(x int, y int) int{
	
	defer func (){
		fmt.Println(recover())
	}()
	//panic("PANIC")
	return x/y

}
func changeYNow(yPtr *int) {
	*yPtr = 100
}
func (rect *Rectange) area() float64 {
	return rect.width * rect.height
}