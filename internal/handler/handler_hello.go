package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/sasd13/go-server-starter/pkg/errors"
)

type HandlerHello struct {}

func NewHandlerHello() *HandlerHello {
	return &HandlerHello{}
}

func (h *HandlerHello) Handle(res http.ResponseWriter, req *http.Request) {
	log.Info().Msg("Handle hello")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(http.StatusOK)

	body, err := json.Marshal(map[string]interface{}{
		"data": "Hello, World!",
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
