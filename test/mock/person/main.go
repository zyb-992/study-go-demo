package person

type Person interface {
	Eat(food string) string
	Sleep(name string) string
}

type Student struct {
	p    Person
	Name string
}

func (p *Student) Eat(food string) string {
	return p.p.Eat(food)
}

func (p *Student) Sleep() string {
	return p.p.Sleep(p.Name)
}
