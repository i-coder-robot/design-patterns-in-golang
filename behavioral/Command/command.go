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
	return Command{person: p, method: method}
}

func (c *Command) Execute() {
	c.method()
}

func NewPerson(name string, cmd Command) Person {
	return Person{name, cmd}
}

func (p *Person) Talk() {
	fmt.Println(fmt.Sprintf("%s is talking.\n", p.name))
	p.cmd.Execute()
}

func (p *Person) PassOn() {
	fmt.Println(fmt.Sprintf("%s is passing on.\n", p.name))
	p.cmd.Execute()
}

func (p *Person) Gossip() {
	fmt.Println(fmt.Sprintf("%s is gossiping.\n", p.name))
	p.cmd.Execute()
}

func (p *Person) Listen() {
	fmt.Println(fmt.Sprintf("%s is Listening.\n", p.name))
}
