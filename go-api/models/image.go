package models

// Image for modeling data
type Image struct {
	ID string `json:"id"`
	UserID string `json:"userid"`
	OriginalFileName string `json:"originalfilename"`
	FilePath string `json:"filepath"`
	Status string `json:"status"`
}