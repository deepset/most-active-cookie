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

* ./most-active-cookie -f cookie_log.csv -d 2018-12-09
