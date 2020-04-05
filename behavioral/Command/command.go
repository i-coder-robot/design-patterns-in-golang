package Command

import "fmt"

type Person struct {
	name string
	cmd  Command
}
type Command struct {
	person *Person
	method func()
}

func NewCommand(p *Person, method func()) Command {
	return Command{
		person: p,
		method: method,
	}
}
func (c *Command) Execute() {
	c.method()
}

func NewPerson(name string, cmd Command) Person {
	return Person{
		name: name,
		cmd:  cmd,
	}
}

func (p *Person) Buy() {
	fmt.Println(fmt.Sprintf("%s is buying ", p.name))
	p.cmd.Execute()
}

func (p *Person) Cook() {
	fmt.Println(fmt.Sprintf("%s is cooking ", p.name))
	p.cmd.Execute()
}

func (p *Person) Wash() {
	fmt.Println(fmt.Sprintf("%s is washing ", p.name))
	p.cmd.Execute()
}

func (p *Person) Listen() {
	fmt.Println(fmt.Sprintf("%s is Listening ", p.name))
}

func (p *Person) Talk() {
	fmt.Println(fmt.Sprintf("%s is Talking ", p.name))
	p.cmd.Execute()
}
