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
)

func getAllImages(year int) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir("./")
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		if strings.Contains(file.Name(), ".png") && strings.Contains(file.Name(), strconv.Itoa(year)) {
			files = append(files, file.Name())
		}
	}
	sort.Strings(files)
	return files, nil
}

func main() {

	// load layout setting and parse them into an object
	var settings generator.Settings
	// read file
	data, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		fmt.Println("error:", err)
	}

	// generate the calendar pages
	generator.Generate(settings)

	images, err := getAllImages(2021)
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
	pdf.WritePdf("2021.pdf")

}
