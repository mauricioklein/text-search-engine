FROM golang:1.9

WORKDIR /go/src/text-search-engine
COPY . .

RUN [ "go-wrapper", "download" ]
RUN [ "go-wrapper", "install" ]

ENTRYPOINT [ "text-search-engine" ]
