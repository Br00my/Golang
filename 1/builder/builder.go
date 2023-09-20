package builder

type Builder struct {
	position string
	salary int
	address string
}

func (p *Builder) SetPosition(position string) {
	p.position = position
}

func (p *Builder) GetPosition() string {
	return p.position
}

func (p *Builder) SetSalary(salary int) {
	p.salary = salary
}

func (p *Builder) GetSalary() int {
	return p.salary
}

func (p *Builder) SetAddress(address string) {
	p.address = address
}

func (p *Builder) GetAddress() string {
	return p.address
}
