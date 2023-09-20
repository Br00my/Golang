package main

import (
    "1/employee"
    "1/programmer"
    "1/builder"
    "1/director"
    "1/officer"
    "1/guard"
    "fmt"
)

func main() {
    p := &programmer.Programmer{}
    employee.SetPosition(p, "Test")
    fmt.Println("Position %s", employee.GetPosition(p))
}
