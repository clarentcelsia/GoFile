package data

const FOrientation = "P"
const FPaper = "A4"
const FUnit = "mm"
const FFontdir = ""

var FMarginX = 10.
var FMarginY = 10.
var FFontFamily = "Arial"

var FLineBreak = 1.
var FLineWidth = 0.5
var FSPLineY = 2.
var FEnterY = 5.
var FTabX = 5.
var FTBFillColor = map[string]interface{}{
	"Red":   255,
	"Green": 255,
	"Blue":  255,
}

var FHeadings = []string{"No", "Header2", "Header3", "Header4", "Header5"}
var FColW = []float64{10.0, 58.0, 30.0, 30.0, 34.0, 25.0}
var FTBH float64 = 5.
var FContents = [][]string{
	{
		"1", "TEXT2", "TEXT3", "TEXT4", "TEXT5",
	},
	{
		"2", "TEXT2", "TEXT3", "TEXT4", "TEXT5",
	},
}

func CreatePDF() (bytes.Buffer, error) {
	gopdf := gofpdf.New(FOrientation, FUnit, FPaper, FFontdir)
	gopdf.SetMargins(FMarginX, FMarginY, FMarginX)
	gopdf.AddPage()
	pagewidth, pageheight := gopdf.GetPageSize()
	_, _, mRight, mbottom := gopdf.GetMargins()
	areaW := pagewidth - 2*FMarginX

	// TITLE
	gopdf.SetFont(FFontFamily, "", 10.7)
	_, lineHt := gopdf.GetFontSize()
	html := gopdf.HTMLBasicNew()
	html.Write(lineHt+FSPLineY,
		fmt.Sprintf(`<center>TITLE<br>%s<br>%s</center>`, strings.ToUpper("SUBTITLE"), strings.ToUpper("SUBTITLE")))

	// LINE
	x, y := gopdf.GetXY()
	y += FLineBreak
	x2 := pagewidth - mRight
	gopdf.SetLineWidth(FLineWidth)
	gopdf.Line(x, y, x2, y)

	// CONTENT
	gopdf.SetFontSize(10.)
	y = gopdf.GetY() + (0.5 * FEnterY)
	gopdf.SetX(FMarginX)
	gopdf.SetY(y)
	gopdf.Write(lineHt+FSPLineY, "CONTENT1")
	gopdf.SetX(gopdf.GetX() + (0.5 * FTabX))
	offcolonX := gopdf.GetX()
	gopdf.Write(lineHt+FSPLineY, " : ")
	gopdf.SetX(gopdf.GetX())
	gopdf.MultiCell(0.6*areaW-gopdf.GetX(), lineHt+FSPLineY, "ALLEN ZHENG", "", "L", false)
	offsetY := gopdf.GetY()

	gopdf.SetXY(0.6*areaW+(FTabX), y)
	offsetX := gopdf.GetX()
	gopdf.Write(lineHt+FSPLineY, "CONTENTX1")
	gopdf.SetX(gopdf.GetX() + FTabX)
	offcolonX2 := gopdf.GetX()
	gopdf.Write(lineHt+FSPLineY, " : ")
	gopdf.SetX(gopdf.GetX())
	gopdf.Write(lineHt+FSPLineY, "THIS EXAMPLE")
	fmt.Println(gopdf.GetX())

	gopdf.SetXY(FMarginX, offsetY)
	gopdf.Write(lineHt+FSPLineY, "CONTENT2")
	gopdf.SetX(offcolonX)
	gopdf.Write(lineHt+FSPLineY, " : ")
	gopdf.SetX(gopdf.GetX())
	gopdf.MultiCell(0.6*areaW-gopdf.GetX(), lineHt+FSPLineY, "ALLEN ZHENG", "", "L", false)
	offsetY2 := gopdf.GetY()
	gopdf.SetXY(offsetX, offsetY)
	gopdf.Write(lineHt+FSPLineY, "CONTENTX2")
	gopdf.SetX(offcolonX2)
	gopdf.Write(lineHt+FSPLineY, " : ")
	gopdf.SetX(gopdf.GetX())
	gopdf.Write(lineHt+FSPLineY, "09-12-2022")

	gopdf.SetXY(FMarginX, offsetY2)
	gopdf.Write(lineHt+FSPLineY, "CONTENT3")
	gopdf.SetX(offcolonX)
	gopdf.Write(lineHt+FSPLineY, " : ")
	gopdf.SetX(gopdf.GetX())
	gopdf.MultiCell(0.6*areaW-gopdf.GetX(), lineHt+FSPLineY, "JUST EXAMPLE", "", "L", false)

	// TABLE
	gopdf.SetXY(FMarginX, (gopdf.GetY() + FEnterY + 1.5))
	gopdf.SetFont(FFontFamily, "B", 9.)
	gopdf.SetFillColor(FTBFillColor["Red"].(int), FTBFillColor["Green"].(int), FTBFillColor["Blue"].(int))

	for col := 0; col < len(FHeadings); col++ {
		gopdf.CellFormat(FColW[col], FTBH, FHeadings[col], "1", 0, "CM", true, 0, "")
	}
	// Draw tbcontent
	gopdf.Ln(-1)
	gopdf.SetFont(FFontFamily, "", 9.)
	gopdf.SetFillColor(FTBFillColor["Red"].(int), FTBFillColor["Green"].(int), FTBFillColor["Blue"].(int))
	height := 0.0
	exList := FContents
	for row := 0; row < len(exList); row++ {
		list := exList[row]
		currX, currY := gopdf.GetXY()
		x := currX
		for cols := 0; cols < len(FHeadings); cols++ {
			rowdata := list[cols]
			// Check if the word fit into the table based on given width of each col.
			// If the word has more than 1 line, set the height of its column to your set size.
			lines := gopdf.SplitLines([]byte(rowdata), FColW[cols])
			h := float64(len(lines))*(lineHt) + 2*float64(len(lines))
			if h > height {
				height = h
			}
		}

		// PageBreak if row doesnt fit on the page.
		if gopdf.GetY()+height > pageheight-mbottom {
			gopdf.AddPage()
			currY = gopdf.GetY()
		}

		for cols := 0; cols < len(FHeadings); cols++ {
			rowdata := list[cols]
			gopdf.Rect(x, currY, FColW[cols], height-0.2, "")
			gopdf.MultiCell(FColW[cols], lineHt+1.4, rowdata, "", "CB", false)
			x += FColW[cols]
			gopdf.SetXY(x, currY)
		}

		gopdf.SetXY(currX, currY+height)
		height = 0
	}

	// FOOTNOTE
	gopdf.SetXY(FMarginX, gopdf.GetY()+FEnterY)
	gopdf.SetFont(FFontFamily, "I", 9.)
	gopdf.WriteAligned(areaW, lineHt+0.5, "THIS IS JUST EXAMPLE", "L")

	var pdfbyte = new(bytes.Buffer)
	err := gopdf.Output(pdfbyte)
	return *pdfbyte, err
}
