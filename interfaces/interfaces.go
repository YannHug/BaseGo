package interfaces

import "fmt"

type Instrumenter interface {
	Play()
}

type Amplifier interface {
	Connect(ampl string)
}

type Guitar struct {
	Strings int
}

type Piano struct {
	Keys int
}

func (g Guitar) Play() {
	fmt.Printf("Tzouuuuuiiinngg with %d strings\n", g.Strings)
}

func (g Guitar) Connect(ampl string) {
	fmt.Printf("Connecté à %s\n", ampl)
}

func (p Piano) Play() {
	fmt.Printf("Plip plong with %d keys\n", p.Keys)
}

func Main() {
	var inst Instrumenter
	inst = &Guitar{5}
	inst.Play()

	inst = &Piano{80}
	inst.Play()

	g := Guitar{12}
	g.Connect("Marshall")
	g.Play()
}
