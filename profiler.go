package main

import (
	"net/http"
	"net/http/pprof"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", hiHandler)

	// регистрация pprof-обработчиков
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	for i := 0; i < 2; i++ {
		go func() {
			for {

			}
		}()
	}

	http.ListenAndServe(":8088", r)
}
