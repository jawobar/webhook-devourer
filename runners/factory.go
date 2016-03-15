package runners

import (
  "log"
)

func Create(name string, params map[string]string) Runner {
  switch name {
  case "bash":
    return NewBashRunner(params["command"], params["args"])
  default:
    log.Fatal("Unknown handler name: " + name)
  }
  return nil
}
