package main

import (
	"net/http"
)

func init() {
	http.HandleFunc("/debug/refresh", func(w http.ResponseWriter, r *http.Request) {
		go data.biz_refresh()
	})
	http.HandleFunc("/debug/crash", func(w http.ResponseWriter, r *http.Request) {
		log.Fatalln("This is to crash the app on purpose")
	})
}
