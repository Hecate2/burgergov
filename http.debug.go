package main

import (
	"net/http"
)

func init() {
	http.HandleFunc("/debug/refresh", setCorsHeaders(func(w http.ResponseWriter, r *http.Request) {
		go data.biz_refresh()
	}))
}
