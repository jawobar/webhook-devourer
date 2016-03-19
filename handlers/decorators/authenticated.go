package decorators

import (
  "net/http"
  "bitbucket.org/jawobar/webhook-devourer/handlers"
)

type AuthenticatedHandler struct {
  Handler handlers.Handler
  Apikeys []string
}

func (auth AuthenticatedHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  if auth.authenticate(req) {
    auth.Handler.ServeHTTP(res, req)
  } else {
    http.Error(res, "Not Authorized", http.StatusUnauthorized)
  }
}

func (auth AuthenticatedHandler) authenticate(req *http.Request) bool {
  key := req.URL.Query().Get("apikey")

  for _, k := range auth.Apikeys {
    if k == key {
      return true
    }
  }

  return false
}
