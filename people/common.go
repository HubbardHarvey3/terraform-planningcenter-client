package people

import (
	"net/http"
	"net/http/httptest"
)

var mockAppId string = "mock-app-id"
var mockSecret string = "mock-app-secret"

func setupMockServer(responseBody string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(statusCode)
			if responseBody != "" {
				w.Write([]byte(responseBody))
			}
		}))
}
