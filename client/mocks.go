package client

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

var mockMap = make(map[string]*Mock)

type Mock struct {
	URL          string
	HTTPMethod   string
	RespHTTPCode int
	RespHeaders  http.Header
	RespBody     string
}

func AddMockups(mocks ...*Mock) error {
	for _, m := range mocks {
		normalizedUrl, err := getNormalizedUrl(m.URL)
		if err != nil {
			return fmt.Errorf("Error parsing mock with url=%s. Cause: %s", m.URL, err.Error())
		}
		mockMap[m.HTTPMethod+" "+normalizedUrl] = m
	}

	return nil
}

func getNormalizedUrl(urlStr string) (string, error) {
	urlObj, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	result := urlStr

	//sorting query param strings
	if len(urlObj.RawQuery) > 0 {
		result = strings.Replace(urlStr, urlObj.RawQuery, "", 1)

		mk := make([]string, len(urlObj.Query()))
		i := 0
		for k := range urlObj.Query() {
			mk[i] = k
			i++
		}
		sort.Strings(mk)
		for i := 0; i < len(mk); i++ {
			if i+1 < len(mk) {
				result = fmt.Sprintf("%s%s=%s&", result, mk[i], urlObj.Query().Get(mk[i]))
			} else {
				result = fmt.Sprintf("%s%s=%s", result, mk[i], urlObj.Query().Get(mk[i]))
			}
		}
	}
	return result, nil
}
