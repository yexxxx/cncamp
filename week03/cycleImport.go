package week03

import (
	"fmt"
	"github.com/yexxxx/cncamp/week01/cycle"
)

var (
	Pack cycle.Pack
)

func init() {
	Pack = cycle.Pack{Name: "yex", Location: "China"}
	println("pkg week02 init")
	fmt.Print("%T", Pack)
}

type User struct {
	Name   string
	Gender bool
	Age    int
}
