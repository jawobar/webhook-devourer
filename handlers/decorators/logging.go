package decorators

import (
  "bytes"
  "log"
  "io/ioutil"
  "net/http"
  "bitbucket.org/jawobar/webhook-devourer/handlers"
)

type LoggingHandler struct {
  Handler handlers.Handler
}

func (logger LoggingHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    log.Fatalf("ERROR: %s", err)
  }

  log.Printf("Received %s from %s\n%s\n", req.Method, req.RemoteAddr, body)
  req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

  logger.Handler.ServeHTTP(res, req)
}
