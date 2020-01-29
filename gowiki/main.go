package gowiki

import (
	"fmt"
	"io/ioutil"
)

const srcPath = "../src/"

type Page struct {
	Title string
	Body []byte
}

func (p *Page) Save() error  {
	filename := p.Title + ".html"
	return ioutil.WriteFile(srcPath + filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(srcPath + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, err
}

func Execute() {
	p1 := &Page{Title: "Test Page", Body: []byte("This is a sample page.")}
	p1.Save()
	p2, _ := LoadPage("Test Page")
	fmt.Println(string(p2.Body))
}