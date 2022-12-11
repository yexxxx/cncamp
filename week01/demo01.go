package main

import (
	"fmt"
	"math/rand"
)

var (
	str1 = "hello world"
)

func main() {
//	testArray()
    testDefer()
}

func convertType() {
	// 表达式T(v)将值v转换为类型T
}

// 由同一个数组或切片构造的切片，创建的是当前数组的引用，修改数组元素会修改切片，反之依然
// append函数会重新开辟空间创建切片
func testArray() {
	var arr1 [10]int

	arr1 = initArrWithRandomInt(arr1)

	var slice1 = arr1[:]
	fmt.Println("init value")
	fmt.Printf("arr1 : %v \n", arr1)
	fmt.Printf("slice1 : %v \n", slice1)

	//modify arr1
	arr1[2] = -1
	fmt.Println("modify arr1[2]=-1")
	fmt.Printf("arr1 : %v \n", arr1)
	fmt.Printf("slice1 : %v \n", slice1)

	//modiy slice1
	slice1[0] = -1
	fmt.Println("modify slice1[0]=-1")
	fmt.Printf("arr1 : %v \n", arr1)
	fmt.Printf("slice1 : %v \n", slice1)

	slice2 := rand.Perm(6)
	slice3 := append(slice1, slice2...)
	fmt.Println("init value")
	fmt.Printf("slice2 : %v \n", slice2)
	fmt.Printf("slice3 : %v \n", slice3)

	//modify slice2
	slice2[0] = -2
	fmt.Println("modify slice2[0]=-2")
	fmt.Printf("slice2 : %v \n", slice2)
	fmt.Printf("slice3 : %v \n", slice3)

	slice1[0] = -2
	fmt.Println("modify slice1[0]=-2")
	fmt.Printf("slice2 : %v \n", slice2)
	fmt.Printf("slice3 : %v \n", slice3)

	slice4 := append(arr1[:], slice2...)
	slice5 := slice4[:]
	fmt.Println("init value")
	fmt.Printf("slice4 : %v \n", slice4)
	fmt.Printf("slice5 : %v \n", slice5)

	slice4[10] = -11
	fmt.Println("modify slice[10]=-11")
	fmt.Printf("slice4 : %v \n", slice4)
	fmt.Printf("slice5 : %v \n", slice5)

	arr1[0] = -10
	fmt.Println("modify arr1[0]=-10")
	fmt.Printf("slice1 : %v \n", slice1)
	fmt.Printf("slice4 : %v \n", slice4)

	fmt.Printf("%p,%p \n", &arr1, &slice1)

}

func initArrWithRandomInt(arr [10]int) (result [10]int) {
	for i, _ := range arr {
		arr[i] = rand.Intn(20)
	}
	return arr
}

func testDefer()  {
    defer testRecover1()

    runtimeError()

    //子协程的异常捕获
    go func() {
        defer testRecover1()
        runtimeError()
    }()


    fmt.Println("recoverd")
//    panic("panic")
}

func testRecover1()  {
    if err := recover(); err != nil {
        fmt.Println(err)
    }
}

func testRecover2() func() {
    return func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }
}

func runtimeError() {
    var div int
    result := 1 / div
    println(result)
}
