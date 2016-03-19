package decorators

import (
  "bytes"
  "log"
  "io/ioutil"
  "net/http"
)

type LoggingHandler struct {
  Handler http.Handler
}

func (logger LoggingHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    log.Fatalf("Error reading request body: %s", err)
  }

  log.Printf("Received %s from %s\n%s\n", req.Method, req.RemoteAddr, body)
  req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

  logger.Handler.ServeHTTP(res, req)
}
