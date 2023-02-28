package main

import (
    "fmt"
    "log"
    "go-api/db"
    "go-api/handler"
    "go-api/models"
    "net/http"
    "os"
)

func main() {
	log.Print("The is Server Running on localhost port 3000")

	// hardcoded test data
	db.Imagedb["001"] = models.Image{ImageID: "001", OriginalFileName: "J_Chow_220914_6835_1200.jpg", FilePath: "savedImages/upload-3983864146.png", Status: "unprocessed"}
    db.Imagedb["002"] = models.Image{ImageID: "002", OriginalFileName: "J_Chow_220914_6846_1200.jpg", FilePath: "savedImages/upload-2154451405.png", Status: "unprocessed"}

    // route goes here

    // test route
	http.HandleFunc("/", handler.TestHandler)
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
