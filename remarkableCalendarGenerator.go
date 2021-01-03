package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"florianuhde.com/remarkableCalendarGenerator/generator"
	"github.com/signintech/gopdf"
	"muzzammil.xyz/jsonc"
)

func getAllImages(year int) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir("./")
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		if strings.Contains(file.Name(), ".png") && strings.Contains(file.Name(), strconv.Itoa(year)+"_") {
			files = append(files, file.Name())
		}
	}
	sort.Strings(files)
	return files, nil
}

func loadSettings() generator.Settings {
	var settings generator.Settings
	data, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		fmt.Print(err)
	}
	jc := jsonc.ToJSON(data) // Calling jsonc.ToJSON() to convert JSONC to JSON
	err = json.Unmarshal(jc, &settings)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(settings)
	return settings
}

func main() {

	// load layout setting and parse them into an object
	settings := loadSettings()
	// read file

	// generate the calendar pages
	generator.Generate(settings)

	images, err := getAllImages(settings.Year)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(images)

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	for i := 0; i < len(images); i++ {
		pdf.AddPage()
		pdf.Image(images[i], 0, 0, gopdf.PageSizeA4)
	}
	pdf.WritePdf(strconv.Itoa(settings.Year) + ".pdf")

}
