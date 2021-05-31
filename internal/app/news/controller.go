package news

import "net/http"

func Index(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello!"))
}
