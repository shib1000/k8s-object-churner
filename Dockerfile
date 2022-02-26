FROM golang:1.17-alpine as builder
MAINTAINER shib1000

WORKDIR /usr/src/app

EXPOSE 8080
#ENV GIN_MODE=release
#
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify && \
    apk add --update curl && \
    curl -o aws-iam-authenticator https://amazon-eks.s3.us-west-2.amazonaws.com/1.21.2/2021-07-05/bin/linux/amd64/aws-iam-authenticator &&\
    cp aws-iam-authenticator /usr/bin/ && chmod +x /usr/bin/aws-iam-authenticator

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /usr/local/bin/koc cmd/k8s-object-churner/main.go

# start from scratch
FROM scratch As final
#FROM gcr.io/distroless/static AS final
WORKDIR /usr/src/app
# Copy our static executable
COPY --from=builder /usr/local/bin/koc  /usr/local/bin/koc
COPY --from=builder /usr/src/app/configs /usr/src/app/configs
COPY --from=builder /usr/src/app/bin/kubeconfig-1.yaml /usr/src/app/bin/kubeconfig.yaml
COPY --from=builder /usr/bin/aws-iam-authenticator /usr/bin/aws-iam-authenticator
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/usr/local/bin/koc"]

