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
    RepoUrl string `json:"repo_url"`
    RepoName string `json:"repo_name"`
    Status string
    StarCount int `json:"star_count"`
  }
}

type DockerHubHandler struct {
  runners []runners.Runner
}

func NewDockerHubHandler(runners ...runners.Runner) *DockerHubHandler {
  return &DockerHubHandler{runners: runners}
}

func (handler DockerHubHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  var message DockerHubMessage

  decoder := json.NewDecoder(req.Body)
  err := decoder.Decode(&message)
  context := handler.populateContext(&message)

  if err != nil {
    log.Print(err)
    http.Error(res, "Could not decode JSON", http.StatusBadRequest)
  } else {
    for _, runner := range handler.runners {
      go runner.Run(context)
    }
  }
}

func (handler DockerHubHandler) populateContext(message *DockerHubMessage) map[string]string {
  context := make(map[string]string)

  context["$TAG$"] = message.PushData.Tag
  context["$PUSHER$"] = message.PushData.Pusher
  context["$REPO$"] = message.Repository.RepoName

  return context
}
