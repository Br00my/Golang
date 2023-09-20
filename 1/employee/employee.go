package employee

type Employee interface {
    SetPosition(position string)
    GetPosition() string

    SetSalary(salary int)
    GetSalary() int

    SetAddress(address string)
    GetAddress() string
}

func SetPosition(e Employee, position string) {
    e.SetPosition(position)
}

func GetPosition(e Employee) string {
    return e.GetPosition()
}

func SetSalary(e Employee, salary int) {
    e.SetSalary(salary)
}

func GetSalary(e Employee) int {
    return e.GetSalary()
}

func SetAddress(e Employee, address string) {
    e.SetAddress(address)
}

func GetAddress(e Employee) string {
    return e.GetAddress()
}
