package main

import "fmt"

type people struct {
	name string
}

type humans interface {
	talk()
}

func (p *people) talk() {
	fmt.Println(p.name, "is talking")
}

func main() {
	raul := people{name: "Raul"}

	talksomething(&raul)
}

func talksomething(h humans) {
	h.talk()
}
