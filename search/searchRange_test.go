package search

import (
	"os"
	"testing"

	Error "github.com/deepset/most-active-cookie/errors"
	"github.com/deepset/most-active-cookie/logger"
)

func init() {
	os.Chdir("../")
	os.Remove("logs.log")
	defer os.Chdir("search/")

	logger.StartLogger()
	logger.InfoLogger.Println("Search Range Testing")
}

func TestGetCookiesByDate_Success(t *testing.T) {

	date := "2018-12-09"
	testCookieList := [][]string{
		{"AtY0laUfhglK3lC7", "2018-12-09"},
		{"SAZuXPGUrfbcn5UA", "2018-12-09"},
		{"5UAVanZf6UtGyKVS", "2018-12-09"},
		{"AtY0laUfhglK3lC7", "2018-12-09"},
		{"SAZuXPGUrfbcn5UA", "2018-12-08"},
		{"4sMM2LxV07bPJzwf", "2018-12-08"},
		{"fbcn5UAVanZf6UtG", "2018-12-08"},
		{"4sMM2LxV07bPJzwf", "2018-12-07"},
	}

	searchObj := New(testCookieList, date)

	//Expected Result
	expectedList := [][]string{
		{"AtY0laUfhglK3lC7", "2018-12-09"},
		{"SAZuXPGUrfbcn5UA", "2018-12-09"},
		{"5UAVanZf6UtGyKVS", "2018-12-09"},
		{"AtY0laUfhglK3lC7", "2018-12-09"},
	}
	expectedLength := len(expectedList)

	//Got result
	gotList, _ := searchObj.GetCookiesByDate()
	gotLength := len(gotList)

	if gotLength != expectedLength {
		t.Errorf("Expected list %v got %v", expectedList, gotList)
		t.FailNow()
	}

}

func TestGetCookiesByDate_DateNotFound(t *testing.T) {

	date := "2018-12-11"
	testCookieList := [][]string{
		{"AtY0laUfhglK3lC7", "2018-12-09"},
		{"SAZuXPGUrfbcn5UA", "2018-12-09"},
		{"5UAVanZf6UtGyKVS", "2018-12-09"},
		{"AtY0laUfhglK3lC7", "2018-12-09"},
		{"SAZuXPGUrfbcn5UA", "2018-12-08"},
		{"4sMM2LxV07bPJzwf", "2018-12-08"},
		{"fbcn5UAVanZf6UtG", "2018-12-08"},
		{"4sMM2LxV07bPJzwf", "2018-12-07"},
	}

	expectedError := Error.ErrDateNotFound

	searchObj := New(testCookieList, date)
	_, err := searchObj.GetCookiesByDate()

	if expectedError == err {
		t.Logf("GetCookiesByDate failed with Expected error %q", err)
	} else {
		t.Errorf("expected error %q got %v", expectedError, err)
		t.FailNow()
	}

}
