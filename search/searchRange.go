package search

import (
	"time"

	Error "github.com/deepset/most-active-cookie/errors"
	"github.com/deepset/most-active-cookie/logger"
)

const (
	DATE_FORMAT = "2006-01-02"
)

type Search interface {
	GetCookiesByDate() ([][]string, error)
}

type Data struct {
	CookieList [][]string
	Date       string
}

func New(cookieList [][]string, date string) *Data {
	f := &Data{cookieList, date}
	return f
}

//GetCookiesByDate will give slice of [{cookies,date}] for the given date using binary search
func (s *Data) GetCookiesByDate() ([][]string, error) {

	logger.InfoLogger.Println("start GetCookiesByDate")

	//get first position and last position of given date in cookie list using binary search
	first, last, err := s.searchRange()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	logger.InfoLogger.Println("Search Range Cookie List : ", s.CookieList[first:last+1])
	logger.InfoLogger.Println("end GetCookiesByDate")

	return s.CookieList[first : last+1], nil

}

//searchRange will get get the first occuranne and last occurance of date using binary search
func (s *Data) searchRange() (int, int, error) {
	logger.InfoLogger.Println("start searchRange")
	if len(s.CookieList) == 0 {
		logger.ErrorLogger.Println(Error.ErrEmptyList)
		return -1, -1, Error.ErrEmptyList
	}
	//first occurrence of date
	start := s.first()
	logger.InfoLogger.Println("Start index in Cookie List : ", start)
	//date not found in cookie list
	if start == -1 {
		logger.ErrorLogger.Println(Error.ErrDateNotFound)
		return -1, -1, Error.ErrDateNotFound
	}
	//last occurance of date
	end := s.last()
	logger.InfoLogger.Println("End index in Cookie List : ", end)
	logger.InfoLogger.Println("end searchRange")
	return start, end, nil
}

//first will return first occurance of given date in cookielist using binary search
func (s *Data) first() int {
	logger.InfoLogger.Println("start first")
	n := len(s.CookieList)
	low, high := 0, n-1
	date := parseDate(s.Date)

	for low <= high {
		mid := (low + high) / 2
		if date.Equal(parseDate(s.CookieList[mid][1])) {
			if mid-1 >= 0 && date.Equal(parseDate(s.CookieList[mid-1][1])) {
				high = mid - 1
			} else {
				return mid
			}
		} else if date.Before(parseDate(s.CookieList[mid][1])) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	logger.InfoLogger.Println("end first")
	return -1
}

//last will return first occurance of given date in cookielist using binary search
func (s *Data) last() int {
	logger.InfoLogger.Println("start last")
	n := len(s.CookieList)
	low, high := 0, n-1
	date := parseDate(s.Date)

	for low <= high {
		mid := (low + high) / 2
		if date.Equal(parseDate(s.CookieList[mid][1])) {
			if mid+1 < n && date.Equal(parseDate(s.CookieList[mid+1][1])) {
				high = mid - 1
			} else {
				return mid
			}
		} else if date.Before(parseDate(s.CookieList[mid][1])) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	logger.InfoLogger.Println("end last")
	return -1
}

//parseDate parses the string to date in yyyy-mm-dd format
func parseDate(date string) time.Time {
	t, _ := time.Parse(DATE_FORMAT, date)
	return t
}
