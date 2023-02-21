package main

import (
	"html/template"
	"bufio"
	"log"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
	"github.com/google/uuid"
)

type Page struct {
	Title string
	Body  []byte
	List  []string
	ImageList []string
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

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()
	fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }

    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

	originalFileName := handler.Filename

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("savedImages", "upload-*.png")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    // return that we have successfully uploaded our file!
    fmt.Println("Successfully Uploaded File")

    path := tempFile.Name()

    // add file path to txt
	f, err := os.OpenFile("filepath.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString("id = " + id + "; Original File Name = " + originalFileName + "; File Path = " + path + "\r\n"); err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
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

	var imagelist []string

	// TODO: Add DB code to save path

	readfilepath, err := os.Open("filepath.txt")
	if err == nil {
		fileScanner := bufio.NewScanner(readfilepath)
 
		fileScanner.Split(bufio.ScanLines)
	  
		for fileScanner.Scan() {
			imagelist = append(imagelist, fileScanner.Text())
		}
	} 

	readfilepath.Close()

	p := &Page{Title: "home", List: list, ImageList: imagelist}

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
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/add/", addHandler)
	http.HandleFunc("/", getRoot)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
