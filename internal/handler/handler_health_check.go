package handler

import (
  "net/http"

  "github.com/rs/zerolog/log"
  resp "github.com/sasd13/go-server-starter/pkg/response"
  "github.com/spf13/viper"
)

type HandlerHealthCheck struct {
  Conf *viper.Viper
}

func NewHandlerHealthCheck(conf *viper.Viper) *HandlerHealthCheck {
  return &HandlerHealthCheck{
    Conf: conf,
  }
}

func (h *HandlerHealthCheck) Handle(res http.ResponseWriter, req *http.Request) {
  log.Info().Msg("Handle health check")
  log.Trace().Msgf("Configuration: %v", h.Conf.AllSettings())
  resp.WriteJSON(
    http.StatusOK,
    map[string]interface{}{
      "status": "ok",
    },
    res,
  )
}
