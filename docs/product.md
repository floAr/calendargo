# Calendago

> Create high quality, personalized ready-to-print calendars easily

## Vision

Calendago will enable users to generate read-to-print customized and personalized calendars.

## User Stories

* I as a user must be able to specify the year I want to generate calendar for and select at least 12 pictures and upload them to the platform. Once I click Generate the system should process my request and generate a PDF with a calendar with the pictures inserted somewhere in the design.

* I as a User I should be able to specify which image should be used for which month. If I specify a particular picture for December, for example, it must not appear on another month on the resulting design.

* I as a User should be able to specify a list of birthdates by inputting pairs of (name, birthdate) this should then highlight the dates on the resulting calendar design.

* I as a User should be able to specify whether I want Public Holidays to be highlighted on the calendar, this should also require me to specify which Country or Set of Holidays/

## Architecture & Technical Design

Calendago is developed as an API-first product in order to open possibility of creating mobile apps in addition to the web frontend we will create. 

We will also experiment with making it a Serverless service which can be used on platforms like AWS Lambda.

### Tech Stack

- Go 1.18+
- Vue 3
- BoltDB or SQLite for storage
- OpenAPI 3.1.x for the API specification

### Configuration

Calendago will be configured via command-line flags when starting the program or Environment variables.

The following concept environment variables, subject to change

```shell
#/bin/sh

# .env 
# Hostname or interface to bind to
CALENDAGO_HOST="localhost"
# Port to bind to 
CALENDAGO_PORT="8000"
# Maximum file size for image uploads
CALENDAGO_MAX_FILE_SIZE="10m"
# The database file path for either SQLite or BoltDB 
CALENDAGO_DB_FILE="./calendago.db"
# Domains to allow Cross-Origin Requests from (for the UI)
CALENDAGO_CORS_DOMAINS="calendago.com"
# Working directory where files are stored and generated
CALENDAGO_WORK_DIR="/tmp/calendago/"
# Schedule for how frequently to clean the working directory (remove files from disk)
CALENDAGO_CLEAN_WORK_DIR_AFTER="2h"
```


## ROADMAP

See [GitHub issues](https://github.com/zikani03/calendago/issues)

### Unscheduled + Ideas

* Combine with other apis to automatically populate the calender with
   * Events from online calendar
   * Nice things to fill the blanks (quotes / images)
* Add more configuration options (Color, different fonts for header and body)
* Build a more stable layouting system (% margins)
* Automatically push to RM2
