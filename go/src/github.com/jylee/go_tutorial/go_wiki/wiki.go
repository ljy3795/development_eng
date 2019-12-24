package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Defining the Data structure for WIKI
type Page struct {
	Title string
	Body  []byte //byte slice (it is expected by io library (not string))
}

// For persistent storage
// -> save the Page's body to textfile
// -> return Error (if success, return nil)
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) // _, _, read-write permission for current user only
}

// Load pages
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, p *Page) {
	t, _ := template.ParseFiles("./template/" + tmpl + ".html")
	t.Execute(w, p)
}

// View Handler (for WIKI Page)
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	// [with template]
	// t, _ := template.ParseFiles("./template/view.html")
	// t.Execute(w, p)
	renderTemplate(w, r, "view", p)
}

// Edit Handler
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title} // If error, make new page
	}
	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	// 	"<form action=\"/save/%s\" method=\"POST\">"+
	// 	"<textarea name=\"body\">%s</textarea><br>"+
	// 	"<input type=\"submit\" value=\"Save\">"+
	// 	"</form>",
	// 	p.Title, p.Title, p.Body)

	// [with template]
	// t, _ := template.ParseFiles("./template/edit.html")
	// // fmt.Println("%T\n", t)
	// t.Execute(w, p)

	renderTemplate(w, r, "edit", p)
}

// main function
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	fmt.Printf("%T\n", p1)
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", viewHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
