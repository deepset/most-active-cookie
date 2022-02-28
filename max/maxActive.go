package max

import (
	Error "github.com/deepset/most-active-cookie/errors"
	"github.com/deepset/most-active-cookie/logger"
)

type ActiveCookie interface {
	MostActiveCookieList() ([]string, error)
}

type Data struct {
	CookieList [][]string
	cookieFreq map[string]int
}

func New(list [][]string) Data {
	return Data{CookieList: list}
}

//MostActiveCookieList gets most active cookies list
func (c *Data) MostActiveCookieList() ([]string, error) {
	logger.InfoLogger.Println("start MostActiveCookieList")
	err := c.getCookieFreqMap()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}
	cookieList := c.getMostActiveCookieList()
	logger.InfoLogger.Println("Most Active Cookie List : ", cookieList)
	logger.InfoLogger.Println("end MostActiveCookieList")
	return cookieList, nil
}

//getCookieFreqMap converts cookie list to frequency map
func (c *Data) getCookieFreqMap() error {
	logger.InfoLogger.Println("start getCookieFreqMap")
	if len(c.CookieList) == 0 {
		logger.ErrorLogger.Println(Error.ErrEmptyList)
		return Error.ErrEmptyList
	}
	c.cookieFreq = make(map[string]int)
	for _, val := range c.CookieList {
		c.cookieFreq[val[0]] += 1
	}
	logger.InfoLogger.Println("Cookie Freq Map : ", c.cookieFreq)
	logger.InfoLogger.Println("end getCookieFreqMap")
	return nil
}

//getMostActiveCookieList gets list of most active cookies with the help of stack and cookie frequency map
func (c *Data) getMostActiveCookieList() []string {
	logger.InfoLogger.Println("start getMostActiveCookieList")
	//create stack
	list, top := make([]Record, len(c.CookieList)), -1
	stack := &stack{top, list}

	for cookie, count := range c.cookieFreq {
		newRecord := Record{cookie, count}
		if stack.isEmpty() {
			stack.push(newRecord)
			continue
		}
		topCookie, _ := stack.peek()

		//if top count equal to incoming freq, add to the stack
		if count == topCookie.count {
			stack.push(newRecord)
		} else if count > topCookie.count {
			//set stack to empty and add the new maximum occurrence
			stack.setEmpty()
			stack.push(newRecord)
		}
	}

	cookieList := make([]string, stack.top+1)
	for i := 0; i < len(cookieList); i++ {
		cookieList[i] = stack.list[i].id
	}

	logger.InfoLogger.Println("end getMostActiveCookieList")
	return cookieList
}
