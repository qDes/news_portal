package news

import "net/http"

func Index(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello!"))
}

func Science(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello Science!"))
}

func Politics(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello Politics!"))
}

func Economy(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello Economy!"))
}
