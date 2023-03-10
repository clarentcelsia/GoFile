package main

import (
	cClient "pdf/controllers/client"
	cFile "pdf/controllers/file"
	cPdf "pdf/controllers/pdf"

	"github.com/gin-gonic/gin"
)

func main() {
	// Activate router for server
	router := gin.Default()
	router.GET("/animal", cClient.Animal)
	router.GET("/pdf", cPdf.DownloadPDFFromReader)
	router.GET("/randompdf", cPdf.GenerateGOFPDF)
	router.GET("/readcsv", cPdf.GetCSVFile)
	router.POST("/csv", cPdf.CreateCSVFile)
	router.POST("/upload", cPdf.UploadFormFile)
	router.POST("/dataupload", cClient.GetDataUpload)
	router.GET("/html", cPdf.GenerateHTMl)

	router.GET("/func/imageresize", cFile.ImageResize)
	router.Run(":8082")
}
