package guard

type Guard struct {
  position string
  salary int
  address string
}

func (p *Guard) SetPosition(position string) {
  p.position = position
}

func (p *Guard) GetPosition() string {
  return p.position
}

func (p *Guard) SetSalary(salary int) {
  p.salary = salary
}

func (p *Guard) GetSalary() int {
  return p.salary
}

func (p *Guard) SetAddress(address string) {
  p.address = address
}

func (p *Guard) GetAddress() string {
  return p.address
}
