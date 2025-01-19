package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/sasd13/go-server-starter/pkg/errors"
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
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(http.StatusOK)

	log.Trace().Msgf("Configuration: %v", h.Conf.AllSettings())

	body, err := json.Marshal(map[string]interface{}{
		"status": "ok",
	})
	if err != nil {
		errors.Write(http.StatusInternalServerError, "fail to marshall JSON", err, res)

		return
	}

	if _, err := res.Write(body); err != nil {
		errors.Write(http.StatusInternalServerError, "fail to write response", err, res)

		return
	}
}
