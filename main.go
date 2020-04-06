package main
import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/",func(w http.ResponseWriter,r* http.Request){
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.RUN(":9090")
}
