# Most Active Cookie
This command line application takes a cookie log file and date in UTC format and prints the most active cookie(s) on that date.

The application is built in Golang and uses binary search algorithm

## Install
git clone github.com/deepset/most-active-cookie

cd github.com/deepset/most-active-cookie

go build


## Example
cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00

$ ./most-active-cookie -f cookie_log.csv -d 2018-12-09

Response

There are 2 most Active Cookies on date 2018-12-09

5UAVanZf6UtGyKVS
AtY0laUfhglK3lC7
