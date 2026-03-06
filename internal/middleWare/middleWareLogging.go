package middleWare

import (
	"log"
	"net/http"
)

type LogMiddleware struct {
	next http.Handler
}

func LoggingMiddleware(next http.Handler) http.Handler {
	lm := LogMiddleware{next: next}
	return lm
}

func (lm LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.Method, r.URL, " - started")
	lm.next.ServeHTTP(w, r)
	log.Printf("%s %s %s", r.Method, r.URL, " - finished")
}
