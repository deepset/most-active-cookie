package fileops

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"

	Error "github.com/deepset/most-active-cookie/errors"
	"github.com/deepset/most-active-cookie/logger"
)

const (
	RFC3339_LAYOUT = "2006-01-02T15:04:05"
)

type Data struct {
	Path string
	Date string
}

func New(fileName string, date string) *Data {
	f := &Data{fileName, date}
	return f
}

func (f *Data) ProcessCSVFile() ([][]string, error) {

	logger.InfoLogger.Println("start ProcessCSVFile")
	file, err := os.Open(f.Path)
	if err != nil {
		logger.ErrorLogger.Println(Error.ErrFileNotFound)
		return nil, Error.ErrFileNotFound
	}
	defer file.Close()

	// Get Headers
	var headers, line []string
	var cookieList [][]string

	reader := csv.NewReader(file)
	reader.Comma = ','

	headers, err = reader.Read()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}

	for {
		line, err = reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		cookie, timestamp, err := processLine(headers, line)
		if err != nil {
			logger.InfoLogger.Printf("Line: %sError: %s\n", line, err)
			log.Printf("Line: %sError: %s\n", line, err)
			continue
		}
		cookieList = append(cookieList, []string{cookie, timestamp})
	}

	logger.InfoLogger.Println("Cookie List from file operation ", cookieList)

	logger.InfoLogger.Println("exit ProcessCSVFile")
	return cookieList, nil
}

func processLine(headers []string, data []string) (string, string, error) {
	logger.InfoLogger.Println("start processLine")
	if len(data) != len(headers) {
		logger.ErrorLogger.Println(Error.ErrDataFormatMismatch)
		return "", "", Error.ErrDataFormatMismatch
	}
	cookie, timestamp := data[0], data[1]

	//validate date - get date in format yyyy-mm-dd
	finalDate, err := validateTimestamp(timestamp)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return "", "", err
	}
	logger.InfoLogger.Println("end processLine")
	return cookie, finalDate, nil
}

func validateTimestamp(date string) (string, error) {
	logger.InfoLogger.Println("start validateTimestamp")
	var t time.Time
	var err error
	t, err = time.Parse(RFC3339_LAYOUT, date[:19])
	if err != nil {
		logger.ErrorLogger.Println(err)
		return "", err
	}
	logger.InfoLogger.Println("end validateTimestamp")
	return t.UTC().String()[:10], nil
}
