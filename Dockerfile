FROM golang:1.23 AS build

WORKDIR /stats

COPY go.mod go.sum ./
COPY cmd/ cmd/.
COPY internal/ internal/.
COPY Makefile .

RUN make deps
RUN make build-release

EXPOSE 8080
ENTRYPOINT ["/stats/bin/web-stats"]

