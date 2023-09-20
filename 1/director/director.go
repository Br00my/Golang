package director

type Director struct {
  position string
  salary int
  address string
}

func (p *Director) SetPosition(position string) {
  p.position = position
}

func (p *Director) GetPosition() string {
  return p.position
}

func (p *Director) SetSalary(salary int) {
  p.salary = salary
}

func (p *Director) GetSalary() int {
  return p.salary
}

func (p *Director) SetAddress(address string) {
  p.address = address
}

func (p *Director) GetAddress() string {
  return p.address
}
