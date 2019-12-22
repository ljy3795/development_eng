package main

import (
	"io/ioutil"
)

// Defining the Data structure for WIKI
type Page struct {
	Title string
	Body  []byte //byte slice (it is expected by io library (not string))
}

// For persistent storage
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
