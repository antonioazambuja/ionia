#!/bin/bash

export CDN_VERSION=10.14.1
export CDN_HOST=ddragon.leagueoflegends.com
export CDN_LANGUAGE=pt_BR
export API_KEY=RGAPI-5adac329-9c25-489f-bf96-5973da9b54a0
export ENDPOINT_REGION=https://br1.api.riotgames.com
export HEADER_API_KEY=X-Riot-Token
export REDIS_URL=localhost
export REDIS_PWD=password
export REDIS_PORT=6379
# export CGO_ENABLED=0 && go test test/v1/*.go -v
go build *.go
./main