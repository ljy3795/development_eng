package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"errors"
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

// validate titles
// edit|view|save와 나머지를 구분해서 slice 형태로 return
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$") 

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	// fmt.Println(r.URL.Path)
	// fmt.Println(m)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second 
}

// load template files at once!
var templates = template.Must(template.ParseFiles("template/edit.html", "template/view.html"))

func renderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, p *Page) {
	// t, err := template.ParseFiles("./template/" + tmpl + ".html")
	err := templates.ExecuteTemplate(w, tmpl+".html", p) // ParseFiles + Execute
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w, "Hello Error", http.StatusInternalServerError)
		return
	}
	// err = t.Execute(w, p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

// View Handler (for WIKI Page)
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/view/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	// [with template]
	// t, _ := template.ParseFiles("./template/view.html")
	// t.Execute(w, p)
	renderTemplate(w, r, "view", p)
}

// Edit Handler
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/edit/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

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

// Save Handler
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/save/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	body := r.FormValue("body") // get values from "NAME='body'" in the form of HTML
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// wrapper function takes three handlers and return http.HandlerFunc
//  http.HandlerFunc is suitable to be passed to the function http.HandleFunc

// CLOSURE인 이유
//  --> the variable fn is enclosed by the closure
//  --> extract title & validate 
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w,r)
			return
		}
		// valid title then the enclosed handler function fn will be called
		fn(w, r, m[2])
	}
}


// main function
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	fmt.Printf("%T\n", p1)
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	// http.HandleFunc("/save/", viewHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
