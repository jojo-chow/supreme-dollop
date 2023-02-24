package main

import (
	"html/template"
	"bufio"
	"log"
	"net/http"
	"fmt"
	"os"
	"io"
	"github.com/google/uuid"
)

type Page struct {
	Title string
	Body  []byte
	ImageList []string
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
    tempFile, err := os.CreateTemp("savedImages", "upload-*.png")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := io.ReadAll(file)
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

var templates = template.Must(template.ParseFiles("index.html"))

func getRoot(w http.ResponseWriter, r *http.Request) {
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

	p := &Page{Title: "home", ImageList: imagelist}

	renderTemplate(w, "index", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/", getRoot)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
