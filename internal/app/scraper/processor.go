package scraper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"news_portal/internal"
	"time"
	"unicode/utf8"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var cache = make(map[string]int)

func ProcessLoop(roc <-chan *internal.RawPage, soc chan<- *internal.RawPage) {
	for page := range roc {
		fmt.Println(page.URL)
		if checkUrl(page.URL) {
			bodyB, err := downloadURL(page.URL)
			if err != nil {
				zap.L().Error("load error", zap.Error(err))
			} else {
				page.HTML = byteDecoder(bodyB)
			}
			time.Sleep(time.Duration(1) * time.Second)
			soc <- page
		}

	}
}

func checkUrl(url string) bool {
	//fmt.Println("cache size", len(cache))
	if _, ok := cache[url]; ok {
		return false
	} else {
		cache[url] = 1
	}
	return true
}

func byteDecoder(byteInput []byte) string {
	result := string(byteInput)
	if !utf8.ValidString(result) {
		v := make([]rune, 0, len(result))
		for i, r := range result {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(result[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		result = string(v)
	}

	return result

}

func downloadURL(url string) ([]byte, error) {
	var (
		//req       *http.Request
		response  *http.Response
		err       error
		bodyBytes []byte
		//status    int
	)

	//client := MakeClient(proxy, timeoutSec)
	client := http.Client{}
	//req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("User-Agent", browser.Random())

	response, err = client.Get(url)

	if err != nil {
		return nil, errors.Wrap(err, "error during http request execution")
	}
	defer response.Body.Close()

	bodyBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error during response body read")
	}
	bodyBytes = bytes.Trim(bodyBytes, "\x00")
	return bodyBytes, nil

}
