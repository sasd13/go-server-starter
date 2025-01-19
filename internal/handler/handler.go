package handler

import (
  "net/http"

  "go.uber.org/fx"
)

type Handler interface {
  Handle(res http.ResponseWriter, req *http.Request)
}

var Module = fx.Options(
  fx.Provide(
    NewHandlerHealthCheck,
    NewHandlerHello,
  ),
)
