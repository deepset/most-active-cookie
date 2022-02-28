package max

import (
	"os"
	"strconv"
	"testing"

	Error "github.com/deepset/most-active-cookie/errors"
	"github.com/deepset/most-active-cookie/logger"
)

func init() {
	os.Chdir("../")
	os.Remove("logs.log")
	defer os.Chdir("max/")

	logger.StartLogger()
	logger.InfoLogger.Println("Max Active Testing")
}

func TestMostActiveCookieList_Success(t *testing.T) {

	list1 := [][]string{{"AtY0laUfhglK3lC7", "2018-12-09"}, {"SAZuXPGUrfbcn5UA", "2018-12-09"}, {"5UAVanZf6UtGyKVS", "2018-12-09"}, {"5UAVanZf6UtGyKVS", "2018-12-09"}, {"AtY0laUfhglK3lC7", "2018-12-09"}}
	list2 := [][]string{{"AtY0laUfhglK3lC7", "2018-12-09"}, {"SAZuXPGUrfbcn5UA", "2018-12-09"}, {"5UAVanZf6UtGyKVS", "2018-12-09"}, {"AtY0laUfhglK3lC7", "2018-12-09"}}
	list3 := [][]string{{"SAZuXPGUrfbcn5UA", "2018-12-09"}, {"5UAVanZf6UtGyKVS", "2018-12-09"}, {"AtY0laUfhglK3lC7", "2018-12-09"}}
	want1 := []string{"AtY0laUfhglK3lC7", "5UAVanZf6UtGyKVS"}
	want2 := []string{"AtY0laUfhglK3lC7"}
	want3 := []string{"SAZuXPGUrfbcn5UA", "5UAVanZf6UtGyKVS", "AtY0laUfhglK3lC7"}

	maxObj1 := New(list1)
	maxObj2 := New(list2)
	maxObj3 := New(list3)

	tests := []struct {
		cookieList Data
		want       []string
	}{
		{maxObj1, want1},
		{maxObj2, want2},
		{maxObj3, want3},
	}

	for key, tc := range tests {
		testName := strconv.Itoa(key + 1)

		t.Run(testName, func(t *testing.T) {

			got, _ := tc.cookieList.MostActiveCookieList()

			if len(got) != len(tc.want) {
				t.Errorf("Got %v want %v", got, tc.want)

			} else {
				t.Logf("Cookies with maximum occurrence : %s", got)
			}

		})

	}
}

func TestMostActiveCookieList_EmptyCookieList(t *testing.T) {

	list1 := [][]string{}
	maxObj1 := New(list1)

	_, err := maxObj1.MostActiveCookieList()

	if err == Error.ErrEmptyList {
		t.Logf("MostActiveCookieList failed with Expected error %q", err)
	} else {
		t.Errorf("expecting error %q, got %q", Error.ErrEmptyList, err)
		t.FailNow()
	}

}
