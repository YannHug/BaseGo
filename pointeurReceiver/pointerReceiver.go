package pointeurReceiver

import (
	"fmt"
	"strings"
)

type Post struct {
	Title     string
	Text      string
	published bool
}

func (p Post) Headline() string {
	return fmt.Sprintf("%v - %v", p.Title, p.Text[:50])
}

func (p *Post) Published() bool {
	return p.published
}

func (p *Post) Publish() {
	p.published = true
}

func (p *Post) UnPublish() {
	p.published = false
}

func UpperTitle(p *Post) {
	p.Title = strings.ToUpper(p.Title)
}

func Main() {
	p := Post{
		Title: "Go release",
		Text: `Go est un langage de programmation compilé et concurrent inspiré de C et Pascal. 
Il a été développé par Google5 à partir d’un concept initial de Robert Griesemer (en), Rob Pike et Ken Thompson.`,
	}
	fmt.Println(p.Headline())
	fmt.Printf("Post publish : %v\n", p.Published())
	p.Publish()
	fmt.Printf("Post publish : %v\n", p.Published())
	p.UnPublish()
	fmt.Printf("Post publish : %v\n", p.Published())

	UpperTitle(&p)
	fmt.Println(p.Headline())

	pythonPost := &Post{
		Title: "Python Intro",
		Text:  "Ceci est du texte python qui doit etre vachement long sinon j'ai une erreur dans le slice du texte lors de l'appel de headline",
	}

	UpperTitle(pythonPost)
	fmt.Println(pythonPost.Headline())
}
