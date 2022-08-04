package models

import "mime/multipart"

type (
	File struct {
		Filename string
		Url      string
	}
	DataUpload struct {
		Animal Animal                `form:"animal"`
		Doc    *multipart.FileHeader `form:"file"`
	}
)
