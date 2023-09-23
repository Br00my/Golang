package main

import (
    "1/employee"
    "1/programmer"
    "1/builder"
    "1/director"
    "1/officer"
    "1/guard"
)

func main() {
    p := &programmer.Programmer{}
    employee.SetPosition(p, "programmer")
    employee.SetSalary(p, 40000)
    employee.SetAddress(p, "Home 2")

    b := &builder.Builder{}
    employee.SetPosition(b, "builder")
    employee.SetSalary(b, 50000)
    employee.SetAddress(b, "Home 1")

    d := &director.Director{}
    employee.SetPosition(d, "director")
    employee.SetSalary(d, 30000)
    employee.SetAddress(d, "Home 3")

    o := &officer.Officer{}
    employee.SetPosition(o, "officer")
    employee.SetSalary(o, 60000)
    employee.SetAddress(o, "Home 4")

    g := &guard.Guard{}
    employee.SetPosition(g, "guard")
    employee.SetSalary(g, 10000)
    employee.SetAddress(g, "Home 5")
}
