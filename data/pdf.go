package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	m "pdf/models"
	"strconv"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
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

func CreatePDF(m pdf.Maroto) pdf.Maroto {
	m.RegisterHeader(func() {
		m.Row(RowHeight, func() {
			m.Text("ANIMAL", props.Text{
				Align: consts.Center,
				Size:  HeaderTxtSize,
				Color: TextColor,
				Style: consts.Bold,
			})
		})
	})

	m.Row(SpaceHeight, func() {})
	//CONTENT
	m.SetBorder(true)
	headings := Headings
	contents := GenerateContents()
	m.TableList(headings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      HeaderTxtSize,
			GridSizes: []uint{1, 2, 2, 7},
		},
		ContentProp: props.TableListContent{
			Size:      TbContentSIze,
			GridSizes: []uint{1, 2, 2, 7},
		},
		HeaderContentSpace:     0.01,
		Line:                   false,
		Align:                  consts.Center,
		VerticalContentPadding: 1,
	})

	m.SetBorder(false)

	ColWidth := ColWidthMax / 3.
	m.Row(SpaceHeight, func() {})
	m.Col(uint(ColWidth), func() {
		m.Col(uint(RowHeight-1), func() {
			m.Text("Owned by")
		})
		m.Col(uint(RowHeight-1), func() {
			m.Text("Designed by")
		})
		m.Col(uint(RowHeight-1), func() {
			m.Text("Written by")
		})
	})
	m.Row(5, func() {})
	m.Col(uint(ColWidth), func() {
		m.Col(uint(RowHeight-1), func() {
			m.Text("Owner")
		})
		m.Col(uint(RowHeight-1), func() {
			m.Text("Allen")
		})
		m.Col(uint(RowHeight), func() {
			m.Text("Go")
		})
	})

	m.Row(15, func() {})
	m.Col(uint(ColWidth), func() {
		m.Col(uint(RowHeight-1), func() {
			m.Text("(                                    )")
		})
		m.Col(uint(RowHeight-1), func() {
			m.Text("(                                    )")
		})
		m.Col(uint(RowHeight-1), func() {
			m.Text("(")
			m.ColSpace(3)
			m.Text(")")
		})
	})

	m.Row(SpaceHeight, func() {})

	m.RegisterFooter(func() {
		m.Row(RowHeight, func() {
			m.Text("THANK YOU", props.Text{
				Size:  TextSize,
				Color: TextColor,
				Align: consts.Center,
			})
		})
	})

	return m
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
