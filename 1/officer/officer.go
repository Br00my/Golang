package officer

type Officer struct {
  position string
  salary int
  address string
}

func (p *Officer) SetPosition(position string) {
  p.position = position
}

func (p *Officer) GetPosition() string {
  return p.position
}

func (p *Officer) SetSalary(salary int) {
  p.salary = salary
}

func (p *Officer) GetSalary() int {
  return p.salary
}

func (p *Officer) SetAddress(address string) {
  p.address = address
}

func (p *Officer) GetAddress() string {
  return p.address
}
