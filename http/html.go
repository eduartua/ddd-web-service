package http

import (
	app "github.com/eduartua/ddd-web-service"
)

func htmlUserHandler(us app.UserStore) *UserHandler {
	uh := UserHandler{
		userStore: us,
	}
	return &uh
}

func htmlStaticHandler() *StaticHandler {
	sh := StaticHandler{}
	return &sh
}

