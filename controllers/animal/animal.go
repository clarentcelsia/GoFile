package controllers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"pdf/data"

	"github.com/gin-gonic/gin"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func Animal(c *gin.Context) {
	//Generate PDF
	GenerateAnimalPDF(c)
}

func GenerateAnimalPDF(c *gin.Context) {
	m := pdf.NewMaroto(data.Orientation, data.Paper)
	m.SetPageMargins(10, 10, 10)
	mpdf := data.CreatePDF(m)

	// Generate PDF to local
	// if err := mpdf.OutputFileAndClose("animal.pdf"); err != nil {
	// 	u.HandleResponse(c, err, "Failed to download file", http.StatusInternalServerError)
	// 	return
	// }
	//=========================================

	// Generate PDF in Byte
	bytePdf, err := mpdf.Output()
	if err != nil {
		fmt.Print("Failed to generate pdf file")
		return
	}

	// save to dummy
	data.BytePDF = bytePdf

	// encrypt bytes to string
	encodeddatabytes := base64.StdEncoding.EncodeToString(data.BytePDF.Bytes())

	var body = map[string]string{
		"data": encodeddatabytes,
	}

	c.JSON(http.StatusOK, body)
}
