package server

import (
  "fmt"
  "net/http"

  "github.com/rs/zerolog/log"
  "github.com/spf13/viper"
  "go.uber.org/fx"
)

type Server struct {
  Conf *viper.Viper
}

func NewServer(conf *viper.Viper) *Server {
  return &Server{
    Conf: conf,
  }
}

func Start(s *Server) {
	port := s.Conf.GetString("LISTEN_PORT")

	log.Info().Msgf("Start server on 0.0.0.0:%s", port)

	http.HandleFunc("/health", s.HandleHealthCheck)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal().Msgf("Fail to start server: %s",
			"err", port, err,
		)
	}
}

var Module = fx.Options(
  fx.Provide(
    NewServer,
  ),
)
