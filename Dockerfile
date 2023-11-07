FROM golang:1.21-alpine as builder

RUN mkdir -p /go/src/github.com/keshu12345/truecaller
WORKDIR /go/src/github.com/truecaller
COPY  .  .
RUN apk add --no-cache gcc musl-dev
RUN go build -o maching-prefixes

FROM alpine:edge
WORKDIR /go/src/github.com/keshu12345/truecaller
COPY --from=builder /go/src/github.com/keshu12345/truecaller .
COPY --from=builder /go/src/github.com/keshu12345/truecaller/server ./server
COPY --from=builder /go/src/github.com/keshu12345/truecaller/config ./config
COPY --from=builder /go/src/github.com/keshu12345/truecaller/toolkit ./toolkit
COPY ./entrypoint.sh .
# REST Service
EXPOSE 8080

ENTRYPOINT ["/bin/ash","/go/src/github.com/truecaller/entrypoint.sh"]