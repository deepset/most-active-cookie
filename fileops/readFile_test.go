package fileops

import (
	"os"
	"testing"

	Error "github.com/deepset/most-active-cookie/errors"
	"github.com/deepset/most-active-cookie/logger"
)

func init() {
	os.Chdir("../")
	os.Remove("logs.log")
	defer os.Chdir("fileops/")

	logger.StartLogger()
	logger.InfoLogger.Println("File ops Testing")
}

func TestProcessCSVFile_Success(t *testing.T) {
	path, date := "./ut-files/cookie_log.csv", "2018-12-09"
	wantLength := 8
	fileObj := New(path, date)
	cookieList, _ := fileObj.ProcessCSVFile()

	gotLength := len(cookieList)

	if gotLength != wantLength {
		t.Errorf("expecting list %v , got %v", cookieList, cookieList)
		t.FailNow()
	}
}

func TestProcessCSVFile_FileNotFound(t *testing.T) {
	path, date := "filenotfound_log.csv", "2018-12-09"
	fileObj := New(path, date)
	expectedError := Error.ErrFileNotFound
	_, err := fileObj.ProcessCSVFile()

	if expectedError == err {
		t.Logf("ProcessCSVFile failed with Expected error %q", err)
	} else {
		t.Errorf("expected error %q got %v", expectedError, err)
		t.FailNow()
	}

}
