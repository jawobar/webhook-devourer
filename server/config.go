package server

var BuildTime string
var AppVersion string
var CommitHash string

type RunnerConfig map[string]string

type HandlerConfig struct {
  Type string
  Auth bool
  Runners map[string]RunnerConfig
}

type ServerConfig struct {
  Handlers map[string]HandlerConfig
  Tls struct {
    Key string
    Cert string
  }
  Apikeys []string
}
