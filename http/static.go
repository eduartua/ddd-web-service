package http

import "net/http"

type StaticHandler struct {
}

func (sh *StaticHandler) Home(w http.ResponseWriter, r *http.Request) {
	v := sh.getHome(r)
	v.Render(w, r, nil)
}