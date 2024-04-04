FROM golang:1.22-alpine as go

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build cmd/main.go

FROM ubuntu:22.04 as postgres

RUN apt-get update
RUN apt-get install -y lsb-release 
RUN apt-get clean all

RUN apt-get install -y wget
RUN apt-get install -y gnupg

RUN sh -c 'echo "deb https://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get -y install postgresql

FROM postgres as final
ENV ENVIRONMENT=production

WORKDIR /app
COPY --from=go /app/main .
COPY --from=go /app/entrypoint.sh .

CMD ["./entrypoint.sh"]