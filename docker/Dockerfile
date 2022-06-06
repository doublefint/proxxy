# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds

FROM golang:1.17 AS build

WORKDIR /go/app

COPY Makefile .
COPY go.mod .
COPY go.sum .
COPY cmd/proxxy/ cmd/proxxy/
COPY internal/ internal/

RUN make install
RUN make build

FROM gcr.io/distroless/base AS distroless

COPY --from=build /go/app/bin/app .

ENTRYPOINT ["./app"]
