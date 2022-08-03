package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
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
