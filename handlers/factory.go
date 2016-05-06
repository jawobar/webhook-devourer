package handlers

import (
  "log"
  "net/http"
  "bitbucket.org/jawobar/webhook-devourer/runners"
)

func Create(name string, runners ...runners.Runner) http.Handler {
  switch name {
  case "dockerhub":
    return NewDockerHubHandler(runners...)
  case "bitbucket":
    return NewBitbucketHandler(runners...)
  default:
    log.Fatal("Unknown handler name: " + name)
  }
  return nil
}
