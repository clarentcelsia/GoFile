package data

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	m "pdf/models"
	"strconv"
)

func ReadCSV() {
	curr_dir, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(curr_dir, "file.csv"))

	defer file.Close()

	// Read each records of the file
	contents, _ := csv.NewReader(file).ReadAll()
	for _, v := range contents {
		fmt.Println(v)
	}
}

func CreateCSV() {
	curr_dir, _ := os.Getwd()
	file, _ := os.Create(filepath.Join(curr_dir, "file.csv"))
	defer file.Close()

	//write to csv
	w := csv.NewWriter(file)
	for i, v := range Contents {
		w.Write(GetCSVContent(i, v))
	}
	defer w.Flush() //writes are 'buffered' call flush to ensure that the content has already written to underlying writer

	// OR
	// w.WriteAll(CSVContents)
}

func GenerateHTML(tmpls []string) {
	templates := template.Must(template.New("").Funcs(template.FuncMap{
		"subtr": subtr,
		"list":  list,
	}).ParseFiles(tmpls...))

	var p bytes.Buffer
	templates.ExecuteTemplate(&p, "page", Contents)

	result := "static/index.html"
	curr_dir, _ := os.Getwd()
	f, _ := os.Create(filepath.Join(curr_dir, result))
	defer f.Close()

	w := bufio.NewWriter(f) //will write to f
	w.WriteString(string(p.Bytes()))
	w.Flush()
}

func subtr(a, b float64) float64 {
	return a - b
}

func list(data ...float64) []float64 {
	return data
}

func GetCSVContent(i int, v m.Animal) []string {
	var csvcontents []string
	csvcontents = append(csvcontents, strconv.Itoa(i+1))
	csvcontents = append(csvcontents, v.Name)
	csvcontents = append(csvcontents, v.Type)
	csvcontents = append(csvcontents, v.Description)

	return csvcontents
}

func GetContent(i int, v m.Animal) []string {
	var contents []string
	contents = append(contents, strconv.Itoa(i+1))
	contents = append(contents, v.Name)
	contents = append(contents, v.Type)
	contents = append(contents, v.Description)

	return contents
}

func GenerateContents() [][]string {
	content := Contents

	var contents_ [][]string
	for i, v := range content {
		contents := GetContent(i, v)
		contents_ = append(contents_, contents)
	}
	return contents_
}
