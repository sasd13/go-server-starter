package server

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/sasd13/go-server-starter/pkg/errors"
)

func (s *Server) HandleHealthCheck(res http.ResponseWriter, req *http.Request) {
	log.Info().Msg("Handle health check")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(http.StatusOK)

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
