package handlers

import (
  "log"
  "encoding/json"
  "net/http"
  "bitbucket.org/jawobar/webhook-devourer/runners"
)

type DockerHubMessage struct {
  CallbackUrl string `json:"callback_url"`
  PushData struct  {
    Images []string
    Pusher string
    PuashedAt int `json:"pushed_at"`
    Tag string
  } `json:"push_data"`
  Repository struct {
    CommentCount int `json:"comment_count"`
    DateCreated int `json:"date_created"`
    Description string
    FullDescription string `json:"full_description"`
    IsOfficial bool `json:"is_official"`
    IsPrivate bool `json:"is_private"`
    IsTrusted bool `json:"is_trusted"`
    Name string
    Namespace string
    Owner string
    Repo_url string `json:"repo_url"`
    Repo_name string `json:"repo_name"`
    Status string
    Star_count int `json:"star_count"`
  }
}

type DockerHubHandler struct {
  runners []runners.Runner
}

func (handler DockerHubHandler) New(runners ...runners.Runner) *DockerHubHandler {
  return &DockerHubHandler{runners: runners}
}

func (handler DockerHubHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  var message DockerHubMessage

  decoder := json.NewDecoder(req.Body)
  err := decoder.Decode(&message)

  if err != nil {
    log.Print(err)
    http.Error(res, "Could not decode JSON", http.StatusBadRequest)
  } else {
    for _, runner := range handler.runners {
      go runner.Run()
    }
  }
}
