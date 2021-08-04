# go-clicker

![Lint Status](https://github.com/rusito-23/go-clicker/actions/workflows/golangci-lint.yml/badge.svg) 
![Build Status](https://github.com/rusito-23/go-clicker/actions/workflows/build.yml/badge.svg)

Fake social network back-end built in Go

## Dev Setup

-  Get started by installing the following packages:
   - [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
   - [pre-commit](https://pre-commit.com/#installation)
   - The pre-commit tools used in this repo are based in [pre-commit-golang](https://github.com/dnephin/pre-commit-golang)
- Install [docker](https://docs.docker.com/get-docker/)
- Run `docker-compose build && docker-compose up`
- The app should start listening on `localhost:8080`
