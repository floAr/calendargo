# Go Calendar Generator

This tool generates monthly calendar templates to print or use with devices like the reMarkable2.
Given a setting.json the tool outputs a ong file per month, as well as a single pdf containign all twelve months.



Generate templates for each month of the year using gg (https://github.com/fogleman/gg) and also embed events from online calendar sources. 

* Generate a set of 12 templates for a given year
* Push those template to the remarkable (and remove old ones)
* Pull events from (different) online calendar sources
* Embed online events into templates
* Allow to run run update as a cron and update the templates accordingly