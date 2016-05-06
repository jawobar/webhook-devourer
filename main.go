package main

import (
  "log"
  "flag"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "bitbucket.org/jawobar/webhook-devourer/server"
)

var listenAddr = flag.String("listen", "0.0.0.0:4000", "<address>:<port> to listen on")
var configFile = flag.String("config", "", "Location of configuration file")

func main() {
  log.Print("Starting...");

  flag.Parse()
  config := server.ServerConfig{}

  yml, err := ioutil.ReadFile(*configFile)
  if err != nil {
    log.Fatal("Error reading: " + *configFile, err)
  }

  err = yaml.Unmarshal(yml, &config)
  if err != nil {
    log.Fatal("Unmarshalling YAML: ", err)
  }

  err = server.Start(*listenAddr, &config)
  if err != nil {
    log.Fatal("Unable to start the server: ", err)
  }
}
