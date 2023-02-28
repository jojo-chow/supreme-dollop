package main

import (
    "fmt"
    "log"
    "go-api/db"
    "go-api/handler"
    "go-api/models"
    "net/http"
    "os"
	"html/template"
	"bufio"
	"io"
	"github.com/google/uuid"
)

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

	// TODO: Call add handler and write to JSON file
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
	// TODO: Remove once read from JSON is added
	var imagelist []string
	readfilepath, err := os.Open("filepath.txt")
	if err == nil {
		fileScanner := bufio.NewScanner(readfilepath)
 
		fileScanner.Split(bufio.ScanLines)
	  
		for fileScanner.Scan() {
			imagelist = append(imagelist, fileScanner.Text())
		}
	} 

	readfilepath.Close()

	err = templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	log.Print("The is Server Running on localhost port 3000")

	// hardcoded test data
	// TODO: Replace hardcoded data with reading JSON file
	db.Imagedb["001"] = models.Image{ImageID: "001", OriginalFileName: "J_Chow_220914_6835_1200.jpg", FilePath: "savedImages/upload-3983864146.png", Status: "unprocessed"}
    db.Imagedb["002"] = models.Image{ImageID: "002", OriginalFileName: "J_Chow_220914_6846_1200.jpg", FilePath: "savedImages/upload-2154451405.png", Status: "unprocessed"}

    // route goes here

    // test route
	// http.HandleFunc("/", handler.TestHandler)

    // some frontend for testing
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/", getRoot)

	// get images
	http.HandleFunc("/images/", handler.GetImages)
	// get a single image
	// e.g. localhost:3000/image/?imageid=002
	http.HandleFunc("/image/", handler.GetImage)
	// Add image
	http.HandleFunc("/image/add", handler.AddImage)
	// delete image
	http.HandleFunc("/image/delete", handler.DeleteImage)

	// listen port
	err := http.ListenAndServe(":3000", nil)
	// print any server-based error messages
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
}
