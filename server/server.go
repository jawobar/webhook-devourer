package server

import (
  "log"
  "net/http"
  "bitbucket.org/jawobar/webhook-devourer/handlers"
  "bitbucket.org/jawobar/webhook-devourer/runners"
)

type AuthenticatedHandler struct {
  handler handlers.Handler
}

func (auth AuthenticatedHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  if authenticate(req) {
    auth.handler.ServeHTTP(res, req)
  } else {
    http.Error(res, "Not Authorized", http.StatusUnauthorized)
  }
}

func Start(addr string, config *ServerConfig) error {
  log.Printf("Magic happens on %s", addr)

  prepareHandlers(config)
  return http.ListenAndServe(addr, nil)
}

func authenticate(req *http.Request) bool {
  key := req.URL.Query().Get("apikey")
  return key == "secret"
}

func prepareHandlers(config *ServerConfig) {
  for route, cfg := range config.Handlers {
    handler := handlers.Create(cfg.Type, prepareRunners(&cfg)...)
    if cfg.Auth {
      http.Handle(route, AuthenticatedHandler{handler})
    } else {
      http.Handle(route, handler)
    }
  }
}

func prepareRunners(config *HandlerConfig) []runners.Runner{
  var activeRunners []runners.Runner

  for name, cfg := range config.Runners {
    activeRunners = append(activeRunners, runners.Create(name, cfg))
  }

  return activeRunners
}
