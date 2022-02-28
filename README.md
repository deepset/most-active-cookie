# Most Active Cookie
This command line application takes a cookie log file (csv format) and date(UTC format) and prints the most active cookie(s) on that date.

The application is built in Golang and uses Binary Search Algorithm

## Install
* git clone github.com/deepset/most-active-cookie

* cd github.com/deepset/most-active-cookie

* go build


## Example
* Here sample log file used is cookie_log.csv

* Run the following command from CLI with file(-f) and date(-d) flag 

* ./most-active-cookie **-f** cookie_log.csv **-d** 2018-12-09

## Project Structure
* **fileops** is used for file operations
* **search** is used to find search range using binary search
* **max** is used to find most active cookie
* **logger** is used for generating log file without log level functioanlity , log file generated is logs.log
* **errors** is used for various errors
