FROM golang:1.14 AS builder
ENV CGO_ENABLED 0
WORKDIR /src
ADD . .
RUN go build -mod vendor -o /vm-evaluator

FROM acicn/alpine:3.12
COPY --from=builder /vm-evaluator vm-evaluator
CMD ["/vm-evaluator"]