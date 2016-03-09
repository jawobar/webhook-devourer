package server

type RunnerConfig map[string]interface{}

type HandlerConfig struct {
  Type string
  Auth bool
  Runners map[string]RunnerConfig
}

type ServerConfig struct {
  Handlers map[string]HandlerConfig
}
