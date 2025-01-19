# go-server-starter

## Get started

1. Create a `.env` file, using [.env.example](./.env.example) as a reference:

```sh
cp .env.example .env
```

2. Start the development server

```sh
make dev
```

## Environment variables

| **Name**    | **Description** | **Options**                                   | **Default** |
| :---------- | :-------------- | :-------------------------------------------- | :---------- |
| LISTEN_PORT | listen port     |                                               | 8080        |
| LOG_LEVEL   | log level       | trace, debug, info, warn, error, fatal, panic | info        |
