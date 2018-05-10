package pp

import "fmt"

type Iperson interface {
	SetName(name string)
	BeforeOut()
	Out()
}

type Person struct {
	Specific Iperson
	name     string
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) Out() {
	p.BeforeOut()
	fmt.Println(p.name + ":go out...")
}

func (p *Person) BeforeOut() {
	if p.Specific == nil {
		return
	}
	p.Specific.BeforeOut()
}

//==============================================================================================================================

type Boy struct {
	Person
}

func (_ *Boy) BeforeOut() {
	fmt.Println("get up...")
}

//==============================================================================================================================

type Girl struct {
	Person
}

func (_ *Girl) BeforeOut() {
	fmt.Println("get up...")
	fmt.Println("gress up...")
}
