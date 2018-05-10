package main

import "github.com/ironmantony/golearning/design_pattern/template_pattern/pp"

func main() {
	var p *pp.Person = &pp.Person{}

	p.Specific = &pp.Boy{}
	p.SetName("steven")
	p.Out()

	//
	//p.Specific = &Girl{}
	//p.SetName("lulu")
	//p.Out()

}
