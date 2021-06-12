package processor

import (
	"news_portal/internal"
	"strings"

	"github.com/antchfx/htmlquery"
	"go.uber.org/zap"
)

func ProcessNews(pageRaw *internal.RawPage) (internal.NewsPage, string) {
	var pageNews internal.NewsPage
	topicConf := map[int]string{
		0: "politics",
		1: "economy",
		2: "science",
	}

	topic := topicConf[pageRaw.IDSource]
	titleXpath := GetTitleXpath(pageRaw.URL)
	textXpath := GetTextXpath(pageRaw.URL)

	pageNews.Title = ProcXpath(pageRaw.HTML, titleXpath)
	pageNews.Text = ProcXpath(pageRaw.HTML, textXpath)
	pageNews.Dttm = pageRaw.DTTM
	pageNews.Url = pageRaw.URL

	return pageNews, topic
}

func ProcXpath(html, xpath string) string {
	doc, err := htmlquery.Parse(strings.NewReader(string(html)))
	if err != nil {
		zap.L().Error("html err", zap.Error(err))
	}
	list := htmlquery.Find(doc, xpath)
	return htmlquery.InnerText(list[0])
}

func GetTitleXpath(url string) string {
	if strings.Contains(url, "aif.ru") {
		return "//h1"
	} else if strings.Contains(url, "ria.ru") {
		return "//div[@class=\"article__title\"]"
	}
	return ""
}

func GetTextXpath(url string) string {
	if strings.Contains(url, "aif.ru") {
		return "//div[@class=\"article_text\"]"
	} else if strings.Contains(url, "ria.ru") {
		return "//div[@class=\"article__text\"]"
	}
	return ""
}
