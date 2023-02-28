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
	"encoding/json"
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

	// Add to db
	// TODO: add to real database if we are rewriting backend to include a database
	image := models.Image{ImageID: id, OriginalFileName: originalFileName, FilePath: path, Status: "unprocessed"}
	db.Imagedb[id] = image

	// write image data to JSON file
	jsonbytes, _ := json.MarshalIndent(image, "", " ")

	f, err := os.OpenFile("allimages.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	n, err := f.Write(jsonbytes)
	if err != nil {
		fmt.Println(n, err)
	}

	if n, err = f.WriteString("\n"); err != nil {
		fmt.Println(n, err)
	}

	// TODO: Call add handler

	http.Redirect(w, r, "/", http.StatusFound)
}

var templates = template.Must(template.ParseFiles("index.html"))

func getRoot(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// read json and add to db
	// Open our jsonFile
	jsonFile, err := os.Open("allimages.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	// we initialize our Users array
	var savedImages []models.Image

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

	//TODO : fix JSON reading code or use a real database
	json.Unmarshal(byteValue, &savedImages)
	fmt.Print(savedImages)

	for _, savedImage := range savedImages {
		db.Imagedb[savedImage.ImageID] = savedImage
	}

	log.Print("The is Server Running on localhost port 3000")

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
	err = http.ListenAndServe(":3000", nil)
	// print any server-based error messages
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
}
