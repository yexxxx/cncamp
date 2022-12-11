package main

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
)

var (
	str = "hello world"
)

func TestFor(t *testing.T) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%s-", string(str[i]))
	}
}

func TestForRange(t *testing.T) {
	for i, c := range str {
		fmt.Printf("%d : %s \n", i, string(c))
	}
}

func TestSwitch(t *testing.T) {
	fmt.Print(t)
	condtion := 1
	switch condtion {
	case 0:
		fmt.Printf("is %d \n", 0)
	case 1:
		fmt.Printf("is %d \n", 1)
		fallthrough
	case 2:
		fmt.Printf("is %d \n", 2)
	default:
		fmt.Printf("is default %d", condtion)
	}
}

func TestSlice(t *testing.T) {
	a := []int{}
	b := []int{1, 2, 3}
	c := a
	a = append(b, 1)

	fmt.Printf("%t", &a == &c)
}

// 声明 	var map1 map[keytype] valuetype
func TestMap(t *testing.T) {
	var usermap = map[string]string{}

	// usermap = *new(map[string]int)
	usermap["age"] = "18"

	for k, v := range usermap {
		println(k, ';', v)
	}

	value, exist := usermap["desc"]
	if exist {
		println("desc", value)
	} else {
		println("desc is not existed")
	}
}

//import cycle not allowed
//func TestInit(t *testing.T) {
//  user := week03.User{}
//  fmt.Println(user)
//}

func TestCloure(t *testing.T) {
	var div int
	result := 1 / div
	fmt.Print(result)

	//		defer func() {
	//		    if r := recover(); r!=nil{
	//		        println("recovered in FuncX")
	//		    }
	//		}()
}

func TestInterface(t *testing.T) {
	var man Show
	man = Human{Name: "tom"}
	fmt.Print(man.ShowName())
}

func TestInterface2(t *testing.T) {
	ifs := make([]interface{}, 10, 10)

	ifs = append(ifs, Human{Name: "aaa"})

	ifs = append(ifs, "bbb")

	for _, e := range ifs {
		//类型断言
		if v, ok := e.(Show); ok {
			fmt.Print(v.ShowName())
		}
	}
}

func TestReflect(t *testing.T) {
	//    reflect.TypeOf() 返回被检查对象的类型
	//    reflect.ValueOf()  返回被检查对象的值
	myMap := make(map[string]string, 10)
	myMap["a"] = "a1"
	tt := reflect.TypeOf(myMap)
	vv := reflect.ValueOf(myMap)

	fmt.Println(tt)
	fmt.Println(vv)

	h := &Human{Name: "tom"}
	//    h.learn()
	//    h.grow()

	hValue := reflect.ValueOf(*h)
	//    hValue.Method(0).Call(nil)
	hFields := hValue.NumField()
	fmt.Println(hFields)
	hMethods := hValue.NumMethod()
	fmt.Println(hMethods)
	for i := 0; i < hMethods; i++ {
		fmt.Printf("Method %d: %v \n", i, hValue.Method(i))
	}
}

func TestGoroutine(t *testing.T) {
	intChan := make(chan int)
	go tinyFunc("gogogo", intChan)

	for i := range intChan {
		println(i)
	}
}

func TestCloseChannel(t *testing.T) {
	chan1 := make(chan int)
	close(chan1)
	v, e := <-chan1
	println(v)
	println(e)
}

func TestSwitchGoroutine(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go tinyFunc("chan1", chan1)
	go tinyFunc("chan2", chan2)

	for i := 0; i < 200; i++ {
		select {
		case v := <-chan1:
			fmt.Println("chan1 " + strconv.Itoa(v))
		case v := <-chan2:
			fmt.Println("chan2 " + strconv.Itoa(v))
		}
	}

}

func TestSwitchGoroutineWithTicker(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)
	ticker := time.NewTicker(1 * time.Second)

	go tinyFunc("chan1", chan1)
	go tinyFunc("chan2", chan2)

	for _ = range ticker.C {
		select {
		case v := <-chan1:
			fmt.Println("chan1 " + strconv.Itoa(v))
		case v := <-chan2:
			fmt.Println("chan2 " + strconv.Itoa(v))
		case t := <-ticker.C:
			fmt.Println(t)
		}
	}

}

func TestContext(t *testing.T) {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)

	select {
	case <-timeoutCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}
}

func TestProducerAndConsumer(t *testing.T) {
	intChan := make(chan int, 10)

    go producer(intChan)
    consumer(intChan)
}

func producer(intChan chan<- int) {
	defer close(intChan)
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		intChan <- rand.Intn(100)
	}
}

func consumer(intChan <-chan int) {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println(<-intChan)
	}
}

func tinyFunc(msg string, intChan chan int) {
	defer close(intChan)
	for i := 0; i < 10; i++ {
		//		fmt.Println(msg)
		intChan <- rand.Intn(100)
	}
}

func produceChan() chan int {
	intChan := make(chan int)
	for i := 0; i < 10; i++ {
		//        fmt.Println(msg)
		intChan <- rand.Intn(100)
	}
	return intChan
}
