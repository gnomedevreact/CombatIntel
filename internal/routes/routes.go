package routes

import "net/http"

func RegisterRouter(mux *http.ServeMux) {
	staticHandler := http.StripPrefix("/app/", http.FileServer(http.Dir("./static/")))

	mux.Handle("/app/", staticHandler)
}
