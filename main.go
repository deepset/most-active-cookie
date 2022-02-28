package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	Error "github.com/deepset/most-active-cookie/errors"
	"github.com/deepset/most-active-cookie/fileops"
	"github.com/deepset/most-active-cookie/logger"
	"github.com/deepset/most-active-cookie/max"
	"github.com/deepset/most-active-cookie/search"
)

const (
	DATE_FORMAT = "2006-01-02"
)

type InputData struct {
	filename string
	date     string
}

func init() {
	os.Remove("logs.log")
	logger.StartLogger()
}

func main() {

	logger.InfoLogger.Println("start main")

	var err error
	inputData, err := getCommandLineData()
	if err != nil {
		logger.ErrorLogger.Println(err)
		Error.ExitGracefully(err)
	}

	//Create file object
	fileObj := fileops.New(inputData.filename, inputData.date)

	logger.InfoLogger.Println("main start")
	//Read file data into an array
	cookieList, err := fileObj.ProcessCSVFile()
	if err != nil {
		logger.ErrorLogger.Println(err)
		Error.ExitGracefully(err)
	}

	//Create search object
	searchObj := search.New(cookieList, inputData.date)
	rangeList, err := searchObj.GetCookiesByDate()
	if err != nil {
		logger.ErrorLogger.Println(err)
		Error.ExitGracefully(err)
	}

	//Create max object
	maxObj := max.New(rangeList)
	maxCookie, err := maxObj.MostActiveCookieList()
	if err != nil {
		logger.ErrorLogger.Println(err)
		fmt.Println(err)
	}

	//Output max cookie
	fmt.Printf("There are %d most Active Cookies on date %s \n", len(maxCookie), inputData.date)
	for i := range maxCookie {
		fmt.Println(maxCookie[i])
	}

	logger.InfoLogger.Println("exit main")
}

func getCommandLineData() (InputData, error) {

	logger.InfoLogger.Println("start getCommandLineData")
	var err error
	if len(os.Args) < 2 {
		return InputData{}, errors.New("filepath and date argument is required. Use -f flag with filename and -d flag with date")
	}
	//flags for command-line arguments
	var fileName string
	flag.StringVar(&fileName, "f", "", "File name")
	var date string
	flag.StringVar(&date, "d", "", "Date in UTC")
	flag.Parse()

	if date, err = validateInputData(fileName, date); err != nil {
		return InputData{}, err
	}
	logger.InfoLogger.Println("exit getCommandLineData")
	return InputData{fileName, date}, nil
}

func validateInputData(fileName string, date string) (string, error) {

	logger.InfoLogger.Println("start validateInputData")
	//File extension Validation
	if fileExtension := filepath.Ext(fileName); fileExtension != ".csv" {
		return "", Error.ErrInvalidFile
	}

	//Date validation
	//Parse string to date in yyyy-mm-dd format
	t, err := time.Parse(DATE_FORMAT, date)
	if err != nil {
		return "", err
	}
	t.Nanosecond()
	return t.UTC().String()[:10], nil

}
