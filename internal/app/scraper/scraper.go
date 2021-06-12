package scraper

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"news_portal/internal"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"go.uber.org/zap"
)

// URLX type for news urls with xpath
type URLX struct {
	LinkSource   string
	Xpath        string
	IDLinkSource int
}

func MakeClient(proxy string, timeoutSec int) *http.Client {
	var transport *http.Transport

	if proxyURL, err := url.Parse(proxy); err != nil {
		transport = &http.Transport{}
		zap.L().Error("MakeClient error", zap.String("function", "MakeClient"), zap.Error(err))
	} else {
		transport = &http.Transport{Proxy: http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}

	return &http.Client{Transport: transport, Timeout: time.Duration(timeoutSec) * time.Second}
}



func PostProcessLoop(roc <-chan *internal.RawPage) {
	for page := range roc {
		if  len(page.HTML)!= 0 {
			fmt.Println("PostProc:", len(page.HTML))
		} else {
			fmt.Println("sasi")
		}

	}
}

// GetURLX tries to extract urls from webpage
func GetURLX(urlX URLX, client *http.Client) []string {
	var (
		result []string
		href   string
	)

	baseURL := getBaseURL(urlX.LinkSource)
	// TODO: replace wrt http

	resp, err := client.Get(urlX.LinkSource)
	if err != nil {
		zap.L().Error("GetURLX error", zap.String("function", "GetURLX"), zap.Error(err))
		return result
	}

	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)
	doc, err := htmlquery.Parse(strings.NewReader(string(htmlData)))
	if err != nil {
		zap.L().Error("GetURLX parse error", zap.String("function", "GetURLX"), zap.Error(err))

		return result
	}
	// TODO recover panic
	list := htmlquery.Find(doc, urlX.Xpath)
	for _, n := range list {
		href = htmlquery.InnerText(n)
		if !strings.HasPrefix(href, "http") {
			if strings.HasPrefix(href, "/") {
				href = baseURL + href
			} else {
				href = baseURL + "/" + href
			}

		}
		//fmt.Println(href) // output @href value without A element.
		if !Find(&result, href) {
			result = append(result, href)
		}

	}
	return result

}

func ScanLoop(url URLX, soc chan<- *internal.RawPage) {
	var page internal.RawPage
	client := &http.Client{}
	currentTime := time.Now()
	for {
		fetchedUrls := GetURLX(url, client)
		for _, item := range fetchedUrls {
			//fmt.Println(item)
			page = internal.RawPage{
				HTML:     "",
				URL:      item,
				IDSource: url.IDLinkSource,
				DTTM:     currentTime.Format("2006-01-02-15-04"),
			}
			soc <- &page

		}
		time.Sleep(time.Duration(1) * time.Minute)
	}

}

func GetScraperConfig() []URLX {
	var res []URLX
	politicsRia := URLX{
		LinkSource:   "https://ria.ru/politicsRia/",
		Xpath:        "//div[@class=\"list-item__content\"]//a/@href",
		IDLinkSource: 0,
	}
	economyRia := URLX{
		LinkSource:   "https://ria.ru/economyRia/",
		Xpath:        "//div[@class=\"list-item__content\"]//a/@href",
		IDLinkSource: 1,
	}
	scienceRia := URLX{
		LinkSource:   "https://ria.ru/scienceRia/",
		Xpath:        "//div[@class=\"list-item__content\"]//a/@href",
		IDLinkSource: 2,
	}
	politicsAif := URLX{
		LinkSource:   "https://aif.ru/politics",
		Xpath:        "//div[@class=\"box_info\"]//a/@href",
		IDLinkSource: 0,
	}
	economyAif := URLX{
		LinkSource:   "https://aif.ru/money",
		Xpath:        "//div[@class=\"box_info\"]//a/@href",
		IDLinkSource: 1,
	}

	res = append(res, politicsRia)
	res = append(res, scienceRia)
	res = append(res, economyRia)
	res = append(res, politicsAif)
	res = append(res, economyAif)
	return res
}

// Find takes a slice and looks for an element in it. If found it will
// return a bool of false.
func Find(slice *[]string, val string) bool {
	for _, item := range *slice {
		if item == val {
			return true
		}
	}
	return false
}

func getBaseURL(urlNews string) string {
	var result strings.Builder
	u, err := url.Parse(urlNews)
	if err != nil {
		zap.L().Error("getBaseURL url.Parse error", zap.String("url", urlNews), zap.Error(err))
	}
	protocol := strings.Split(urlNews, "/")[0]
	result.WriteString(protocol)
	result.WriteString("//")
	result.WriteString(u.Hostname())

	return result.String()
}
