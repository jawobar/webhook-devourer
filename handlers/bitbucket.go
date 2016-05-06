package handlers

import (
  "log"
  "encoding/json"
  "net/http"
  "bitbucket.org/jawobar/webhook-devourer/runners"
)

type Link struct {
  Href string
}

type ReferenceState struct {
  Type string
  Name string
  Target struct {
    Type string
    Hash string
    Author User
    Message string
    Date string
    Parents []struct {
      Type string
      Hash string
      Links struct {
        Self Link
        Html Link
      }
    }
    Links struct {
      Self Link
      Html Link
    }
  }
  Links struct {
    Self Link
    Commits Link
    Html Link
  }
}

type User struct {
  Username string
  DisplayName string `json:"display_name"`
  Uuid string
  Links struct {
    Self Link
    Html Link
    Avatar Link
  }
}

type BitbucketMessage struct {
  Actor User
  Repository struct {
    Name string
    FullName string `json:"full_name"`
    Uuid string
    Scm string
    IsPrivate bool `json:"is_private"`
    Links struct {
      Self Link
      Html Link
      Avatar Link
    }
  }
  Push struct {
    Changes []struct {
      New ReferenceState
      Old ReferenceState
      Links struct {
        Html Link
        Diff Link
        Commit Link
      }
      Created bool
      Closed bool
      Forced bool
      Truncated bool
      Commits []struct {
        Hash string
        Type string
        Message string
        Author User
        Links struct {
          Self Link
          Html Link
        }
      }
    }
  }
}

type BitbucketHandler struct {
  runners []runners.Runner
}

func NewBitbucketHandler(runners ...runners.Runner) *BitbucketHandler {
  return &BitbucketHandler{runners}
}

func (handler BitbucketHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  var message BitbucketMessage

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

func (handler BitbucketHandler) populateContext(message *BitbucketMessage) map[string]string {
  context := make(map[string]string)

  for _, change := range message.Push.Changes {
    if change.New.Type == "branch" {
      context["$BRANCH$"] = change.New.Name
      break
    }
  }

  context["$PUSHER$"] = message.Actor.Username
  context["$REPO$"] = message.Repository.Name

  return context
}
