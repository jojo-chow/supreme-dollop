package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
	List  []string
}

func (p *Page) save() error {
	filename := "savedData/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}


func loadPage(title string) (*Page, error) {
	filename := "savedData/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}


func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}


func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}


func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	pageName := r.FormValue("pagename")
	http.Redirect(w, r, "/edit/"+pageName, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("index.html", "edit.html", "view.html"))

func getRoot(w http.ResponseWriter, r *http.Request) {
    entries, err := os.ReadDir("./savedData/")
    if err != nil {
        log.Fatal(err)
    }
    
	var list []string

    for _, e := range entries {
		entryFileName := e.Name()
		entryName := entryFileName[:len(entryFileName)-4]
		list = append(list, entryName)
    }

	p := &Page{Title: "home", List: list}

	renderTemplate(w, "index", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
		    http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/add/", addHandler)
	http.HandleFunc("/", getRoot)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
