package runners

import (
  "log"
)

func Create(name string, params map[string]interface{}) Runner {
  switch name {
  case "bash":
    return BashRunner{}.New(params["command"].(string), params["args"].(string))
  default:
    log.Fatal("Unknown handler name: " + name)
  }
  return nil
}
