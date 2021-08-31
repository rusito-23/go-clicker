# :mouse: go-clicker

![CI Status](https://github.com/rusito-23/go-clicker/actions/workflows/continous_integration.yml/badge.svg)

Incremental game built in Go.

## :memo: Requirements


<details>
    <summary>Click to expand the dev requirements</summary>

- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
- [pre-commit](https://pre-commit.com/#installation)
- [docker](https://docs.docker.com/get-docker/)
- go 1.16

</details>

## :wrench: Set up & Run

- Run `pre-commit install`
- Run `docker-compose build && docker-compose up`
- The app should start listening on `localhost:8080`

- The [docker file](./Dockerfile) uses [CompileDaemon](github.com/githubnemo/CompileDaemon) to automatically reload the app
- Run `docker exec -ti go-clicker_database_1 psql go-clicker-db -U root` to access the `postgresql` CLI and see the DB contents

## :page_with_curl: API Docs

The app provides the following methods:

| Method | Path | Response Example |
| --- | --- | --- |
| `GET` | `/game/ping` | <pre>{<br/>&nbsp;"message" : "pong"<br/>}</pre> |
| `POST` | `/game` | <pre>{<br/>&nbsp;"game": {<br/>&nbsp;&nbsp;"external_id":"3f7ee89f-5bac-4622-8b60-bca93131e21b",<br/>&nbsp;&nbsp;"created_at":"2021-08-08T20:36:03Z",<br/>&nbsp;&nbsp;"updated_at":"2021-08-30T23:08:23Z",<br/>&nbsp;&nbsp;"click_score":0,<br/>&nbsp;&nbsp;"status":"created"<br/>&nbsp;}<br/>}</pre> |
| `GET` | `/game/:external_id` | <pre>{<br/>&nbsp;"game": {<br/>&nbsp;&nbsp;"external_id":"3f7ee89f-5bac-4622-8b60-bca93131e21b",<br/>&nbsp;&nbsp;"created_at":"2021-08-08T20:36:03Z",<br/>&nbsp;&nbsp;"updated_at":"2021-08-30T23:08:23Z",<br/>&nbsp;&nbsp;"click_score":11,<br/>&nbsp;&nbsp;"status":"started"<br/>&nbsp;}<br/>}</pre> |
| `PUT` | `/game/:external_id/click` | <pre>{<br/>&nbsp;"game": {<br/>&nbsp;&nbsp;"external_id":"3f7ee89f-5bac-4622-8b60-bca93131e21b",<br/>&nbsp;&nbsp;"created_at":"2021-08-08T20:36:03Z",<br/>&nbsp;&nbsp;"updated_at":"2021-08-30T23:08:23Z",<br/>&nbsp;&nbsp;"click_score":11,<br/>&nbsp;&nbsp;"status":"started"<br/>&nbsp;}<br/>}</pre> |
| `DELETE` | `/game/:external_id` | <pre>{<br/>&nbsp;"game": {<br/>&nbsp;&nbsp;"external_id":"3f7ee89f-5bac-4622-8b60-bca93131e21b",<br/>&nbsp;&nbsp;"created_at":"2021-08-08T20:36:03Z",<br/>&nbsp;&nbsp;"updated_at":"2021-08-30T23:08:23Z",<br/>&nbsp;&nbsp;"click_score":11,<br/>&nbsp;&nbsp;"status":"started"<br/>&nbsp;}<br/>}</pre> |
| `GET` | `/game/scoreboard?count=2` | <pre>{<br/>&nbsp;"games":[<br/>&nbsp;&nbsp;{<br/>&nbsp;&nbsp;&nbsp;"updated_at":"2021-08-07T00:38:49Z",<br/>&nbsp;&nbsp;&nbsp;"click_score":0<br/>&nbsp;&nbsp;},{<br/>&nbsp;&nbsp;&nbsp;"updated_at":"2021-08-30T23:08:23Z",<br/>&nbsp;&nbsp;&nbsp;"click_score":11<br/>&nbsp;&nbsp;}<br/>&nbsp;]<br/>}</pre> |
