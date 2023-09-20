package programmer

type Programmer struct {
	position string
	salary int
	address string
}

func (p *Programmer) SetPosition(position string) {
	p.position = position
}

func (p *Programmer) GetPosition() string {
	return p.position
}

func (p *Programmer) SetSalary(salary int) {
	p.salary = salary
}

func (p *Programmer) GetSalary() int {
	return p.salary
}

func (p *Programmer) SetAddress(address string) {
	p.address = address
}

func (p *Programmer) GetAddress() string {
	return p.address
}
