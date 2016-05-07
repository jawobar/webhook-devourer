package runners

import (
  "log"
  "strings"
  "os/exec"
)

type BashRunner struct {
  command string
  args string
}

func NewBashRunner(command string, args string) *BashRunner {
  return &BashRunner{command: command, args: args}
}

func (runner *BashRunner) Run(context map[string]string) {
  name := Eval(runner.command, context)
  params := Eval(runner.args, context)

  log.Println("Running script:", name, params)
  args := strings.Split(params, " ")
  output, err := exec.Command(name, args...).Output()

  if err != nil {
    log.Println("Error executing bash script: " + err.Error())
  }
  if len(output) > 0 {
    log.Print("Script output:\n" + string(output))
  }
}
