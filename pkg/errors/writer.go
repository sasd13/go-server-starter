package errors

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func Write(code int, msg string, err error, res http.ResponseWriter) {
	log.Error().Msgf("%s: %s", msg, err)
	http.Error(res, msg, code)
}
