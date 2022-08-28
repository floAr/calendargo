This tool generates monthly calendar templates to print or use with devices like the reMarkable2.
Given a setting.json the tool outputs a ong file per month, as well as a single pdf containign all twelve months.

*If you are only interested int the images / pdf for 2021 have a look in the [output folder](https://github.com/floAr/CalenderGeneratorGo/tree/main/output)*

## Feature
* Generate monthly calendars in the likes of [example calendar for april](https://github.com/floAr/CalenderGeneratorGo/blob/main/example.png "example calendar for april")
* Automatically fill blank days before and after the month to have a nice box layout and always the same first weekday
* Configure layouts and other things on the fly

## Usage
Pull the repository, edit the `setting.json` if needed and run `go run .\remarkableCalendarGenerator.go`

## Customize

Editing the setting.json allows to specify how the calendar should be layouted. The setting.json shipped with the repository is ment to be used to generate templates for the remarkable2 and is structured as follows:

~~~json
{
    "year" : 2021, 
    "width"  : 1404,
    "height"  : 1872,
    "marginLeft"  : 130,
    "marginRight"  : 10,
    "marginTop"  : 5,
    "marginBottom"  : 200,
    "startOfTheWeek"  : 1,
    "headerFontSize":25,
    "headerFont": "arial"
}
~~~
`year` is used to specify the year, `width` and `height` define the total resoltuion of the genrated images (pdf is always A4 at the moment), the `margin` values specify how many empty pixels should be included on each border, `startOfTheWeek` allows to change the first day of the week (0 => Sunday, 1 =S> Monday...), `headerFontSize`specifies the size of the header font (and also its spacing) and `headerFont` the font to use (which must be installed on your system).

