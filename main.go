package main

import (
	cAnimal "pdf/controllers/animal"
	cPdf "pdf/controllers/pdf"

	"github.com/gin-gonic/gin"
)

func main() {
	// Activate router for server
	router := gin.Default()
	router.GET("/animal", cAnimal.Animal)
	router.GET("/pdf", cPdf.DownloadPDFFromReader)
	router.Run(":8082")
}
