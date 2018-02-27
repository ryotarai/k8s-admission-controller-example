FROM golang:1.10.0-alpine

RUN apk add --no-cache --update alpine-sdk

COPY ./webhook /go/src/github.com/ryotarai/k8s-admission-controller-example/webhook
RUN cd /go/src/github.com/ryotarai/k8s-admission-controller-example/webhook && go install .

FROM alpine:3.4

#RUN apk add --update ca-certificates openssl

COPY --from=0 /go/bin/webhook /usr/local/bin/webhook

ENTRYPOINT ["webhook"]
