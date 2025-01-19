package main

import (
  "github.com/sasd13/go-server-starter/internal/config"
  "github.com/sasd13/go-server-starter/internal/handler"
  "github.com/sasd13/go-server-starter/internal/server"
  "github.com/sasd13/go-server-starter/pkg/logger"
  "github.com/rs/zerolog/log"
  _ "github.com/joho/godotenv/autoload"
  "github.com/spf13/viper"
  "go.uber.org/fx"
)

var conf *viper.Viper

func main() {
  conf = config.NewConfig()
  logger.InitLogger(conf)

  log.Debug().Msg("Init")
  app := fx.New(
    config.Module,
    server.Module,
    handler.Module,
    fx.Invoke(server.Start),
  )
  app.Run()
}
