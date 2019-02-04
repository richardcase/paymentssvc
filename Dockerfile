########## Build ###################

FROM golang:1.11.2-alpine3.8 as builder

ENV PACKAGES="\
    build-base \
    git \
"

ENV GO111MODULE=on

RUN apk update && apk add --no-cache $PACKAGES
RUN wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.13.2

COPY . $GOPATH/src/fithub.com/richardcase/Code/poc/paymentssvc/
WORKDIR $GOPATH/src/fithub.com/richardcase/Code/poc/paymentssvc/

RUN go mod download

RUN make ci && make release


########## Output Image ###################
FROM scratch

COPY --from=builder /go/bin/payments-server /app/payments-server

EXPOSE 8080

ENTRYPOINT ["/app/payments-server"]

