// Temporary save the data
package data

import (
	"bytes"
	m "pdf/models"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
)

const Orientation = consts.Portrait
const Paper = consts.Letter

var TbHeaderSize = float64(10)
var TbContentSIze = float64(9)
var RowHeight = float64(5)
var ColWidthMax = float64(12)
var SpaceHeight = float64(10)
var HeaderTxtSize = float64(11)
var TextSize = float64(9)
var TextColor = color.Color{
	Red:   0,
	Green: 0,
	Blue:  0,
}
var Headings = []string{"No", "Name", "Type", "Description"}
var Contents = []m.Animal{
	{
		Name:        "Bonnie",
		Type:        "Mammal",
		Description: "Bonnie is a mammal with 2.0 inches of nose length",
	},
	{
		Name:        "Katty",
		Type:        "Mammal",
		Description: "Katty is a cute mammal cat",
	},
	{
		Name:        "Piscea",
		Type:        "Fish",
		Description: "Piscea is a moody fish",
	},
}
var BytePDF bytes.Buffer
