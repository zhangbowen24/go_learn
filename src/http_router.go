package http_router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "hello %s!\n", ps.ByName("name"))
}
func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/Hello/:name", Index)
	/*http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}", t)
		w.Write([]byte(timeStr))
	})*/
	log.Fatal(http.ListenAndServe(":8080", router))
}
