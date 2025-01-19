package handler

import (
  "net/http"

  "github.com/rs/zerolog/log"
  resp "github.com/sasd13/go-server-starter/pkg/response"
)

type HandlerHello struct {}

func NewHandlerHello() *HandlerHello {
  return &HandlerHello{}
}

func (h *HandlerHello) Handle(res http.ResponseWriter, req *http.Request) {
  log.Info().Msg("Handle hello")
  resp.WriteJSON(
    http.StatusOK,
    map[string]interface{}{
      "data": "Hello, World!",
    },
    res,
  )
}
