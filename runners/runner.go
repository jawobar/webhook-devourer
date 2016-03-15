package runners

import (
  "regexp"
  "strings"
)

type Runner interface {
  Run(context map[string]string)
}

var tokenRegexp = regexp.MustCompile("\\$\\w*\\$")

func Eval(str string, context map[string]string) string {
  for _, v := range tokenRegexp.FindAllString(str, -1) {
    str = strings.Replace(str, v, context[v], -1)
  }
  return str
}
