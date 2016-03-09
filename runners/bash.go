package runners

import (
  "log"
  "os/exec"
)

type BashRunner struct {
  command string
  args []string
}

func (runner BashRunner) New(command string, args ...string) *BashRunner {
  log.Println("Create New: ", command, args)
  return &BashRunner{command: command, args: args}
}

func (runner BashRunner) Run() {
  log.Println("Running script: ", runner.command, runner.args)
  output, err := exec.Command(runner.command, runner.args...).Output()

  if err != nil {
    log.Println("Error executing bash script")
    log.Println(err)
  } else {
    log.Println("Script output:\n", string(output))
  }
}
