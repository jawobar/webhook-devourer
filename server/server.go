package server

import (
  "log"
  "net/http"
  "bitbucket.org/jawobar/webhook-devourer/handlers"
  "bitbucket.org/jawobar/webhook-devourer/runners"
)

var serverConfig *ServerConfig

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

  serverConfig = config
  prepareHandlers()

  if config.Tls.Key != "" && config.Tls.Cert != "" {
    return http.ListenAndServeTLS(addr, config.Tls.Cert, config.Tls.Key, nil)
  } else {
    return http.ListenAndServe(addr, nil)
  }
}

func authenticate(req *http.Request) bool {
  key := req.URL.Query().Get("apikey")

  for _, k := range serverConfig.Apikeys {
    if k == key {
      return true
    }
  }

  return false
}

func prepareHandlers() {
  for route, cfg := range serverConfig.Handlers {
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
