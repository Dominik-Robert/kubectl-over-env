FROM golang:1.17 as builder
WORKDIR /go/src/github.com/Dominik-Robert/kubectl-over-env/
COPY . ./
RUN go mod download \
  && CGO_ENABLED=0 GOOS=linux go build -o config-parse .

FROM ubuntu AS downloader
WORKDIR /root/
ENV KUBECTL_VERSION=v1.23.5
RUN apt update && apt install -y curl \
  && curl -LO https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl \
  && chmod +x ./kubectl


FROM ubuntu
LABEL AUTHOR=Dominik-Robert
LABEL PROJECT=https://github.com/Dominik-Robert/kubectl-over-env
ADD config_gotemplate /
ADD entrypoint.sh /
COPY --from=builder /go/src/github.com/Dominik-Robert/kubectl-over-env/config-parse /
COPY --from=downloader /root/kubectl /
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]