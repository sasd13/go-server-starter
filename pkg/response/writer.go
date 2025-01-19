package response

import (
  "encoding/json"
  "net/http"

  "github.com/rs/zerolog/log"
)

func WriteJSON(status int, body map[string]interface{}, res http.ResponseWriter) error {
  log.Debug().Msgf("Write JSON response: %d", status)
  res.Header().Set("Content-Type", "application/json; charset=utf-8")
  res.WriteHeader(status)

  resBody, err := json.Marshal(body)
  if err != nil {
    msg := "fail to write JSON response"
    log.Error().Msgf("%s: %s", msg, err)
    WriteError500(msg, err, res)

    return err
  }

  if _, err := res.Write(resBody); err != nil {
    msg := "fail to write response"
    log.Error().Msgf("%s: %s", msg, err)
    WriteError500(msg, err, res)

    return err
  }

  return nil
}
