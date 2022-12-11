package week03

import (
    "fmt"
    "github.com/yexxxx/cncamp/week02"
)

var (
    Pack week02.Pack
)


func init() {
    Pack = week02.Pack{Name: "yex", Location: "China"}
    println("pkg week02 init")
    fmt.Print("%T",Pack)
}

type User struct {
	Name   string
	Gender bool
	Age    int
}
