FROM golang:1.22
LABEL maintainer="osouzaelias@gmail.com"

ENV REGION="us-west-2"
ENV PORT=8080
ENV TABLES="node1,node2,node3"

COPY .. /app

WORKDIR /app/cmd

RUN go build -o go-sharding

CMD ["./go-sharding"]