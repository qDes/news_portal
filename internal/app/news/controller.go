package news

import (
	"context"
	"html/template"
	"net/http"
	feed "news_portal/api/proto"
	"regexp"

	"go.uber.org/zap"
)

var (
	svc = GetSvc()
)
const (
	regex = `<.*?>`
	head = `<html>

    <head>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-giJF6kkoqNQ00vy+HMDP7azOuL0xtbfIcaT9wjKHr8RbDVddVHyTfAAsrekwKmP1" crossorigin="anonymous">

        <title>
            NewsNet
        </title>
    </head>
    <body>

    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <div class="container-fluid">
            <a class="navbar-brand" href="/">News</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" href="/">Home</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" href="/science">Science</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" href="/economy">Economy</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" href="/politics">Politics</a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>`
)
type Render struct {
	Title string
	Text  string
	Date  string
}

func Index(resp http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	res, err := svc.FeedClient.GetFeed(ctx, &feed.GetFeedRequest{Topic: "all"})
	if err != nil {
		zap.L().Error("index handler err on grpc call", zap.Error(err))
	}
	var render []Render
	for _, item := range res.Feed {
		render = append(render, Render{
			Title: item.Title,
			Text:  item.Text,
			Date:  item.Date,
		})
	}

	const tpl = `
	<body>
	<h1>{{.Title}} News</h1>
	{{range $element := .Items}}
		<div class="card" style="width: 100rem;">
		  <div class="card-body">
			<h5 class="card-title">{{$element.Title}}</h5>
			<p class="card-text">{{$element.Text}}</p>

		  </div>
		</div>
	{{end}}
	</body>
</html>`

	t, err := template.New("webpage").Parse(head+tpl)
	if err != nil {
		zap.L().Error("page parse error", zap.Error(err), zap.String("function", "Index"))
	}

	data := struct {
		Title string
		Items []Render
	}{
		Title: "All",
		Items: render,
	}
	err = t.Execute(resp, data)
	if err != nil {
		zap.L().Error("execute error", zap.Error(err), zap.String("function", "Index"))
	}
}

func Science(resp http.ResponseWriter, req *http.Request) {

	ctx := context.Background()
	res, err := svc.FeedClient.GetFeed(ctx, &feed.GetFeedRequest{Topic: "science"})
	if err != nil {
		zap.L().Error("science handler err on grpc call", zap.Error(err))
	}
	var render []Render
	for _, item := range res.Feed {
		render = append(render, Render{
			Title: item.Title,
			Text:  item.Text,
			Date:  item.Date,
		})
	}

	const tpl = `
	<body>
	<h1>{{.Title}} News</h1>
	{{range $element := .Items}}
		<div class="card" style="width: 100rem;">
		  <div class="card-body">
			<h5 class="card-title">{{$element.Title}}</h5>
			<p class="card-text">{{$element.Text}}</p>

		  </div>
		</div>
	{{end}}
	</body>
</html>`

	t, err := template.New("webpage").Parse(head+tpl)
	if err != nil {
		zap.L().Error("page parse error", zap.Error(err), zap.String("function", "Science"))
	}

	data := struct {
		Title string
		Items []Render
	}{
		Title: "Science",
		Items: render,
	}
	err = t.Execute(resp, data)
	if err != nil {
		zap.L().Error("execute error", zap.Error(err), zap.String("function", "Science"))
	}
}

func Politics(resp http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	res, err := svc.FeedClient.GetFeed(ctx, &feed.GetFeedRequest{Topic: "politics"})
	if err != nil {
		zap.L().Error("politics handler err on grpc call", zap.Error(err))
	}
	var render []Render
	for _, item := range res.Feed {
		render = append(render, Render{
			Title: item.Title,
			Text:  item.Text,
			Date:  item.Date,
		})
	}

	const tpl = `
	<body>
	<h1>{{.Title}} News</h1>
	{{range $element := .Items}}
		<div class="card" style="width: 100rem;">
		  <div class="card-body">
			<h5 class="card-title">{{$element.Title}}</h5>
			<p class="card-text">{{$element.Text}}</p>

		  </div>
		</div>
	{{end}}
	</body>
</html>`

	t, err := template.New("webpage").Parse(head+tpl)
	if err != nil {
		zap.L().Error("page parse error", zap.Error(err), zap.String("function", "Politics"))
	}

	data := struct {
		Title string
		Items []Render
	}{
		Title: "Politics",
		Items: render,
	}
	err = t.Execute(resp, data)
	if err != nil {
		zap.L().Error("execute error", zap.Error(err), zap.String("function", "Politics"))
	}
}

func Economy(resp http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	res, err := svc.FeedClient.GetFeed(ctx, &feed.GetFeedRequest{Topic: "economy"})
	if err != nil {
		zap.L().Error("economy handler err on grpc call", zap.Error(err))
	}
	var render []Render
	for _, item := range res.Feed {
		render = append(render, Render{
			Title: item.Title,
			Text:  item.Text,
			Date:  item.Date,
		})
	}

	const tpl = `
	<body>
	<h1>{{.Title}} News</h1>
	{{range $element := .Items}}
		<div class="card" style="width: 100rem;">
		  <div class="card-body">
			<h5 class="card-title">{{$element.Title}}</h5>
			<p class="card-text">{{$element.Text}}</p>

		  </div>
		</div>
	{{end}}
	</body>
</html>`

	t, err := template.New("webpage").Parse(head+tpl)
	if err != nil {
		zap.L().Error("page parse error", zap.Error(err), zap.String("function", "Economy"))
	}

	data := struct {
		Title string
		Items []Render
	}{
		Title: "Economic",
		Items: render,
	}
	err = t.Execute(resp, data)
	if err != nil {
		zap.L().Error("execute error", zap.Error(err), zap.String("function", "Economy"))
	}
}

// This method uses a regular expresion to remove HTML tags.
func stripHtmlRegex(s string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(s, "")
}