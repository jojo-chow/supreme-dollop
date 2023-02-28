package models

// Image for modeling data
type Image struct {
	ImageID string `json:"imageid"`
	UserID string `json:"userid"`
	OriginalFileName string `json:"originalfilename"`
	FilePath string `json:"filepath"`
	Status string `json:"status"`
}