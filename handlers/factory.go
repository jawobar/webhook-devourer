package handlers

import (
  "log"
  "bitbucket.org/jawobar/webhook-devourer/runners"
)

func Create(name string, runners ...runners.Runner) Handler {
  switch name {
  case "dockerhub":
    return DockerHubHandler{}.New(runners...)
  default:
    log.Fatal("Unknown handler name: " + name)
  }
  return nil
}
