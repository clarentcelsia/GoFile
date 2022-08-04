package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"pdf/data"
	m "pdf/models"
	u "pdf/utils"

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

func GetDataUpload(c *gin.Context) {
	var dupload m.DataUpload
	c.ShouldBind(&dupload)

	dir, _ := os.Getwd()
	fmt.Println(dir)
	err := c.SaveUploadedFile(dupload.Doc, (dir + "/" + dupload.Doc.Filename))
	if err != nil {
		u.HandleResponse(c, err, "Failed get the data", http.StatusInternalServerError)
		return
	}
	val, bool := c.GetPostForm("animal")
	if !bool {
		u.HandleResponse(c, nil, "Data animal not found", http.StatusInternalServerError)
		return
	}

	//convert json string to map
	var inter map[string]interface{}
	json.Unmarshal([]byte(val), &inter)
	dupload.Animal = m.Animal{
		Name:        inter["Animal"].(map[string]interface{})["Name"].(string),
		Type:        inter["Animal"].(map[string]interface{})["Type"].(string),
		Description: inter["Animal"].(map[string]interface{})["Description"].(string),
	}

	c.JSON(http.StatusOK, dupload)

	// =========================== OR ===================================
	// mr, errmr := c.Request.MultipartReader()
	// if errmr != nil {
	// 	fmt.Println(errmr)
	// 	return
	// }

	// // loops data in multipart form
	// for {
	// 	fmt.Println("START LOOPING")
	// 	part, errpart := mr.NextPart()
	// 	if errpart == io.EOF {
	// 		break //stop
	// 	}
	// 	if errpart != nil {
	// 		http.Error(c.Writer, errpart.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	if part.FormName() == data.FormFieldName {
	// 		// get current dir
	// 		file, errcreate := os.Create((dir + "/" + part.FileName()))
	// 		if errcreate != nil {
	// 			fmt.Println("Can't create file")
	// 			http.Error(c.Writer, errcreate.Error(), http.StatusInternalServerError)
	// 			return
	// 		}
	// 		defer file.Close()

	// 		// read part then write it to file
	// 		_, errcopy := io.Copy(file, part)
	// 		if errcopy != nil {
	// 			fmt.Println("Can't read and write file")
	// 			return
	// 		}
	// 	}

	// 	var an m.Animal
	// 	if part.FormName() == "animal" {
	// 		jsondecoder := json.NewDecoder(part)
	// 		fmt.Println(jsondecoder)
	// 		errdecode := jsondecoder.Decode(&an)
	// 		if errdecode != nil {
	// 			http.Error(c.Writer, errdecode.Error(), http.StatusInternalServerError)
	// 			return
	// 		})
	// 	}

	// }
}
