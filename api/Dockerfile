FROM golang:1.17.2-alpine3.14 as builder

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod ./
# COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080

COPY ./src ${ROOT}
RUN GOOS=linux go build -o ${ROOT}/binary

#####

FROM scratch as deployer

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 8080
CMD ["/go/src/app/binary"]