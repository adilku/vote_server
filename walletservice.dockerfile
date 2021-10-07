FROM golang:1.16-alpine

WORKDIR /wallet_server/


COPY . ./

RUN ls -la ./*

RUN go mod download -x

ENTRYPOINT ["sh", "./build.sh"]