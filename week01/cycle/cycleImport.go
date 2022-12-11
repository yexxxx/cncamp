package cycle

import (
    "fmt"
    "github.com/yexxxx/cncamp/week03"
)

var (
    User week03.User
)


func init() {
	User = week03.User{Name: "yex", Age: 28}
    println("pkg week02 init")
    fmt.Print("%T",User)
}

type Pack struct {
	Name     string
	Location string
}
