package server

import (
  "fmt"
  "log"
  "net/http"
  "bitbucket.org/jawobar/webhook-devourer/handlers"
  "bitbucket.org/jawobar/webhook-devourer/handlers/decorators"
  "bitbucket.org/jawobar/webhook-devourer/runners"
)

var serverConfig *ServerConfig

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

func prepareHandlers() {
  for route, cfg := range serverConfig.Handlers {
    handler := handlers.Create(cfg.Type, prepareRunners(&cfg)...)

    if cfg.Auth {
      handler = decorators.AuthenticatedHandler{handler, serverConfig.Apikeys}
    }
    if cfg.Log {
      handler = decorators.LoggingHandler{handler}
    }

    http.Handle(route, handler)
  }

  http.HandleFunc("/status", func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Build time:", BuildTime)
    fmt.Fprintln(res, "App version:", AppVersion)
    fmt.Fprintln(res, "Commit hash:", CommitHash)
  })
}

func prepareRunners(config *HandlerConfig) []runners.Runner{
  var activeRunners []runners.Runner

  for name, cfg := range config.Runners {
    activeRunners = append(activeRunners, runners.Create(name, cfg))
  }

  return activeRunners
}
