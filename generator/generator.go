package generator

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/flopp/go-findfont"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/jinzhu/now"
)

// Settings is a structure to pass in every layout related property from outside
type Settings struct {
	Width  int
	Height int

	MarginLeft   int
	MarginRight  int
	MarginTop    int
	MarginBottom int

	HeaderFont     string
	HeaderFontSize int

	CellHeight     float64
	StartOfTheWeek time.Weekday
}

var settings Settings

var cellWidth float64 = 0
var cellHeight float64 = 0
var font *truetype.Font

func drawCell(x float64, y float64, date int, dc *gg.Context) {
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(x, y, cellWidth, cellHeight)
	dc.Stroke()
	if date > 0 {
		dc.DrawStringAnchored(strconv.Itoa(date), x, y, -0.2, 1.2)
	} else {
		dc.DrawRectangle(x, y, cellWidth, cellHeight)
		dc.SetRGB(0.8, 0.8, 0.8)
		dc.Fill()
	}

}

func generateMonthGrid(year int, month time.Month, startDay time.Weekday, dc *gg.Context) {
	// maxRows := 6

	t := time.Date(year, month, 1, 0, 0, 0, 0, time.Now().Location())

	leadingEmptyDays := 0
	for i := 7; i >= 0; i-- {
		if t.AddDate(0, 0, i).Weekday() == startDay {
			break
		}
		leadingEmptyDays++
	}

	// Draw cells
	x := 0
	y := 0
	dayCount := now.With(t).EndOfMonth().Day()
	for i := 0; i < dayCount+leadingEmptyDays; i++ {
		x = i % 7
		y = i / 7
		drawCell(cellWidth*float64(x), cellHeight*float64(y), i+1-leadingEmptyDays, dc)
	}

	for i := x + 1; i < 7; i++ {
		drawCell(cellWidth*float64(i), cellHeight*float64(y), -1, dc)
	}

	dc.Translate(0, cellHeight*6)
}

func generateMonthHeader(year int, month time.Month, dc *gg.Context) float64 {
	var headerString string = month.String() + " " + strconv.Itoa(year)
	_, h := dc.MeasureString(headerString)
	spacing := h * 3
	dc.Translate(0, spacing/2)
	dc.DrawStringAnchored(headerString, float64(dc.Width())/2, 0, 0.5, 1.2)
	dc.Translate(0, spacing)
	return spacing * 1.5

}

func generateWeekHeader(startDay time.Weekday, dc *gg.Context) float64 {
	t := time.Now()
	offset := 0
	for i := 0; i < 7; i++ {
		if t.AddDate(0, 0, i).Weekday() == startDay {
			break
		}
		offset++
	}
	for i := 0; i < 7; i++ {
		dc.DrawStringAnchored(t.AddDate(0, 0, i+offset).Weekday().String(), float64(i)*cellWidth+cellWidth/2, float64(settings.HeaderFontSize), 0.5, 0.5)

		dc.DrawRectangle(float64(i)*cellWidth, 0, cellWidth, float64(settings.HeaderFontSize)*2)
		dc.SetRGB(0, 0, 0)
		dc.Stroke()
	}
	dc.Translate(0, float64(settings.HeaderFontSize)*2)
	return float64(settings.HeaderFontSize) * 2
}

func monthToStringNr(month time.Month) string {
	i := int(month)
	if i >= 10 {
		return strconv.Itoa(i)
	}
	return "0" + strconv.Itoa(i)
}

func generateMonth(year int, month time.Month) {
	// monthly logic
	dc := gg.NewContext(settings.Width, settings.Height)
	cellWidth = float64(settings.Width-settings.MarginLeft-settings.MarginRight) / 7

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)

	face := truetype.NewFace(font, &truetype.Options{Size: 48})
	dc.SetFontFace(face)
	mHeaderHeight := generateMonthHeader(year, month, dc)
	dc.Translate(120, 0) // move to side for menue bar
	face = truetype.NewFace(font, &truetype.Options{Size: float64(settings.HeaderFontSize)})
	dc.SetFontFace(face)
	wHeaderHeight := generateWeekHeader(settings.StartOfTheWeek, dc)

	cellHeight = (float64(settings.Height-settings.MarginTop-settings.MarginBottom) - wHeaderHeight - mHeaderHeight) / 6

	generateMonthGrid(year, month, settings.StartOfTheWeek, dc)

	dc.SavePNG(strconv.Itoa(year) + "_" + monthToStringNr(month) + "_" + month.String() + ".png")
}

// Generate creates a set of calendar templates for the given settings
func Generate(set Settings) {
	settings = set
	year := 2021

	fontPath, err := findfont.Find(settings.HeaderFont + ".ttf")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found '"+settings.HeaderFont+"' in '%s'\n", fontPath)

	// load the font with the freetype library
	fontData, err := ioutil.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	font, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	generateMonth(year, time.January)
	generateMonth(year, time.February)
	generateMonth(year, time.March)
	generateMonth(year, time.April)
	generateMonth(year, time.May)
	generateMonth(year, time.June)
	generateMonth(year, time.July)
	generateMonth(year, time.August)
	generateMonth(year, time.September)
	generateMonth(year, time.October)
	generateMonth(year, time.November)
	generateMonth(year, time.December)
}
