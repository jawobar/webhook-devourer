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
  runner.command = Eval(runner.command, context)
  runner.args = Eval(runner.args, context)

  log.Println("Running script:", runner.command, runner.args)
  args := strings.Split(runner.args, " ")
  output, err := exec.Command(runner.command, args...).Output()

  if err != nil {
    log.Println("Error executing bash script")
    log.Println(err)
  } else {
    log.Print("Script output:\n" + string(output))
  }
}
