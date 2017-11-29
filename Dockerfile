FROM golang:1.8.3-alpine3.6

COPY domain-bot-stackoverflow-poller /root/domain-bot-stackoverflow-poller

WORKDIR /root

CMD ["./domain-bot-stackoverflow-poller"]