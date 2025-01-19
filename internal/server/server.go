package server

import (
  "fmt"
  "net/http"
  
  "github.com/rs/zerolog/log"
  "github.com/sasd13/go-server-starter/internal/handler"
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

func Start(
  s *Server,
  health *handler.HandlerHealthCheck,
  hello *handler.HandlerHello,
) {
	port := s.Conf.GetString("LISTEN_PORT")

	log.Info().Msgf("Start server on 0.0.0.0:%s", port)

	http.HandleFunc("/health", health.Handle)
  http.HandleFunc("/hello", hello.Handle)

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
