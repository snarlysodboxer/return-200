FROM golang:1.8 AS go-build
RUN go get github.com/tools/godep
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/return-200
COPY main.go /go/src/return-200/
COPY Godeps /go/src/return-200/Godeps
COPY vendor /go/src/return-200/vendor
RUN godep go build -o /return-200 && chmod ug+x /return-200

FROM scratch
MAINTAINER david amick <docker@davidamick.com>
COPY --from=go-build /return-200 /return-200
ENTRYPOINT ["/return-200"]
CMD ["-h"]

