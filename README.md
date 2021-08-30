# :mouse: go-clicker

![CI Status](https://github.com/rusito-23/go-clicker/actions/workflows/continous_integration.yml/badge.svg)

Incremental game built in Go.

## :memo: Requirements


<details>
    <summary>Dev Requirements</summary>

- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
- [pre-commit](https://pre-commit.com/#installation)
- [docker](https://docs.docker.com/get-docker/)
- go 1.16

</details>

## :wrench: Set up & Run

- Run `pre-commit install`
- Run `docker-compose build && docker-compose up`
- The app should start listening on `localhost:8080`

A couple of considerations:

- The [docker file](./Dockerfile) uses [CompileDaemon](github.com/githubnemo/CompileDaemon) to automatically reload the app
- Run `docker exec -ti go-clicker_database_1 psql go-clicker-db -U root` to access the `postgresql` cli and see the DB contents

## :tada: Acknowledgments

- The pre-commit tools used in this repo are based in [pre-commit-golang](https://github.com/dnephin/pre-commit-golang)
