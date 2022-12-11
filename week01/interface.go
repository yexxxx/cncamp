package main

import "fmt"

type Show interface {
    ShowName() string
}


type Human struct {
    Name string
    Age int
}

func (h Human) ShowName() string {
    return h.Name
}

func (h Human) Learn() string {
    fmt.Println("learning........")
    return "learning........"

}

//func (h *Human) Grow()  {
//    fmt.Println("grow day by day........")
//}