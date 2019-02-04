[![CircleCI](https://circleci.com/gh/richardcase/paymentssvc.svg?style=svg)](https://circleci.com/gh/richardcase/paymentssvc) [![Coverage Status](https://coveralls.io/repos/github/richardcase/paymentssvc/badge.svg?branch=master)](https://coveralls.io/github/richardcase/paymentssvc?branch=master) ![Go Report Card](https://goreportcard.com/badge/github.com/richardcase/paymentssvc)

# Payments Service Sample

This a sample that implements a rudimentary payments service. It has:

* RESTful API
* CQRS with ES

To learn more about the design go [here](docs/design.md).

> There are some issues with the current implementation and with hindsight go-swagger wouldn't be used and if we were going to use CQRS/ES then we would need to split the service into separate write/read services and probably use Kakfa Streams instead as its more appropriate.

## Running Locally

To run locally you need to run Mongo:

```bash
docker run -d --name mongo -p 27017:27017 mongo:latest
```

Then run and run the payments service:

```bash
make build
./payments-server -scheme http
```

## Ares for future work

The following are areas for future work:

* [ ] Expose metrics (so Prometheus can scrape)
* [ ] Create Kubernetes artefacts for deployment
* [ ] Add CircleCI workflow to build &b publish docker image
* [ ] Authentication & Authorization


