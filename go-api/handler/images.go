package handler

import (
    "go-api/models"
    "go-api/db"
    "go-api/utils"
    "net/http"
    "encoding/json"
)

func TestHandler(res http.ResponseWriter, req *http.Request) {
	// Add the response return message
	HandlerMessage := []byte(`{
	 "success": true,
	 "message": "The server is running properly"
	 }`)
   
	utils.ReturnJsonResponse(res, http.StatusOK, HandlerMessage)

}

// Get Images handler
func GetImages(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
   
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "Check your HTTP method: Invalid HTTP method executed",
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
	 return
	}
   
	var images []models.Image
   
	for _, image := range db.Imagedb {
	 images = append(images, image)
	}
   
	// parse the image data into json format
	imageJSON, err := json.Marshal(&images)
	if err != nil {
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "Error parsing the image data",
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
	 return
	}
   
	utils.ReturnJsonResponse(res, http.StatusOK, imageJSON)
}


// Get a single image handler
func GetImage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "Check your HTTP method: Invalid HTTP method executed",
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
	 return
	}
   
	if _, ok := req.URL.Query()["imageid"]; !ok {
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "This method requires the image id",
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
	 return
	}
   
	id := req.URL.Query()["imageid"][0]
   
	image, ok := db.Imagedb[id]
	if !ok {
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "Requested image not found",
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
	 return
	}
   
	// parse the image data into json format
	imageJSON, err := json.Marshal(&image)
	if err != nil {
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "Error parsing the image data",
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
	 return
	}
   
	utils.ReturnJsonResponse(res, http.StatusOK, imageJSON)
}

// Add a image handler
func AddImage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "Check your HTTP method: Invalid HTTP method executed",
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
	 return
	}
   
	var image models.Image
   
	payload := req.Body
   
	defer req.Body.Close()
	// parse the image data into json format
	err := json.NewDecoder(payload).Decode(&image)
	if err != nil {
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "Error parsing the image data",
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
	 return
	}
   
	db.Imagedb[image.ImageID] = image
	// Add the response return message
	HandlerMessage := []byte(`{
	 "success": true,
	 "message": "Image was successfully created",
	 }`)
   
	utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
}

// Delete a image handler
func DeleteImage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "Check your HTTP method: Invalid HTTP method executed",
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
	 return
	}
   
	if _, ok := req.URL.Query()["imageid"]; !ok {
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "This method requires the image id",
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusBadRequest, HandlerMessage)
	 return
	}
   
	imageid := req.URL.Query()["imageid"][0]
	image, ok := db.Imagedb[imageid]
	if !ok {
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "Requested image not found",
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
	 return
	}
	// parse the image data into json format
	imageJSON, err := json.Marshal(&image)
	if err != nil {
	 // Add the response return message
	 HandlerMessage := []byte(`{
	  "success": false,
	  "message": "Error parsing the image data"
	 }`)
   
	 utils.ReturnJsonResponse(res, http.StatusBadRequest, HandlerMessage)
	 return
	}
   
	utils.ReturnJsonResponse(res, http.StatusOK, imageJSON)
}
