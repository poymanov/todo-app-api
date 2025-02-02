package swagger

import (
	"github.com/swaggo/http-swagger/v2"
	"net/http"
	_ "poymanov/todo/docs"
)

func NewSwaggerHandler(router *http.ServeMux) {
	router.Handle("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8099/swagger/doc.json"),
	))
}
