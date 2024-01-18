package http

import app "github.com/eduartua/ddd-web-service"

func jsonUserHandler(us app.UserStore) *UserHandler {
	uh := UserHandler{
		userStore: us,
	}
	return &uh
}