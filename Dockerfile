FROM golang:1.8 as builder

WORKDIR /go/src/github.com/smpio/kube-yaml-cleaner/
RUN go get -d -v github.com/ghodss/yaml
COPY . .

ARG GOOS=linux
ARG GOARCH=amd64

RUN make -f Makefile.docker


FROM scratch
COPY --from=builder /go/src/github.com/smpio/kube-yaml-cleaner/kube-yaml-cleaner /
ENTRYPOINT ["/kube-yaml-cleaner"]
