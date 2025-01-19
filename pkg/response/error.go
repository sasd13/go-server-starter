package response

import (
  "net/http"
)

func WriteError(status int, msg string, err error, res http.ResponseWriter) {
  http.Error(res, msg, status)
}

func WriteError500(msg string, err error, res http.ResponseWriter) {
  WriteError(http.StatusInternalServerError, msg, err, res)
}
