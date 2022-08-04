package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"pdf/data"
	m "pdf/models"
	u "pdf/utils"
)

func DownloadPDFFromReader(c *gin.Context) {
	response, err := http.Get("http://localhost:8082/animal")
	if err != nil {
		c.Status(http.StatusServiceUnavailable)
	}

	var result map[string]interface{}
	resbody, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(resbody, &result)

	data := result["data"].(string) // encoded file
	databyte, _ := base64.StdEncoding.DecodeString(data)

	// reader := response.Body
	// defer reader.Close()

	headers := map[string]string{
		"Content-Disposition": `attachment; filename="animal.pdf"`,
	}
	reader := bytes.NewReader(databyte)
	c.DataFromReader(http.StatusOK, response.ContentLength, response.Header.Get("Content-Type"), reader, headers)

	// ========================

	// ABLE TO USE THIS WITH C.FILE IF DATA BEEN SAVED TO LOCAL FIRST (OUTPUT AND SAVEFILE)
	// c.Writer.Header().Set("Content-Disposition", `attachment; filename="animal.pdf"`)
	// c.File("animal.pdf")

	//......
	// reader := bytes.NewReader(decodebytefile)
	// _, errread := io.Copy(c.Writer, reader)
	// if errread != nil {
	// 	fmt.Println("Failed to copy file")
	// 	return
	// }

}

// MULTIPART
func UploadFormFile(c *gin.Context) {
	var animal m.Animal
	c.BindJSON(&animal)

	req, w := io.Pipe()

	// writing to pipe writer
	multi := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer multi.Close()

		// get file from local and read file
		pdf, errresponse := os.Open("C:\\Clarenti\\Data\\Project\\Sampling/animal.pdf")
		if errresponse != nil {
			u.HandleResponse(c, errresponse, "Can't Open File", http.StatusInternalServerError)
			return
		}
		defer pdf.Close()

		// *I'm not gonna use this stat since I need the name of the file only.
		// fileinfo, errstat := pdf.Stat()
		// if errstat != nil {
		// 	fmt.Println("Can't see file info")
		// 	return
		// }

		// Multipart File
		iowriter, errcreate := multi.CreateFormFile(data.FormFieldName, pdf.Name())
		if errcreate != nil {
			u.HandleResponse(c, errcreate, "Can't create form file", http.StatusInternalServerError)
			return
		}
		// copy pdf file to writer so it's not gonna read the entire file into memory,
		// especially when reads a large file.
		_, errcopy := io.Copy(iowriter, pdf)
		if errcopy != nil {
			fmt.Println("Can't copy the file")
			return
		}

		// Multipart JSON
		params := map[string]interface{}{
			"Animal": animal,
		}
		bodybytes, _ := json.Marshal(params)
		multi.WriteField("animal", string(bodybytes))

	}()

	// Send data to url
	response, errresponse := http.Post("http://localhost:8082/dataupload", multi.FormDataContentType(), req)
	if errresponse != nil {
		u.HandleResponse(c, errresponse, "Failed to upload file", http.StatusInternalServerError)
		return
	}

	var content = make([]byte, response.ContentLength)
	response.Body.Read(content)
	response.Body.Close()

	var result map[string]interface{}
	json.Unmarshal(content, &result)

	u.HandleResponse(c, result, "Upload succeed", response.StatusCode)

	// ===============================================================
	// SIMULATE (If this func returns *http.request) -> http.NewRequest()
	// client := &http.Client{}
	// response, errresponse := client.Do(request)
	// if errresponse != nil {
	// 	fmt.Println(errresponse)
	// 	return
	// }
}
