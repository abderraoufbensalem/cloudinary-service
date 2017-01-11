FROM golang:latest

ADD . /go/src/github.com/abderraoufbensalem/cloudinary-service
WORKDIR /go/src/github.com/abderraoufbensalem/cloudinary-service
RUN go get github.com/gorilla/mux
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/opentracing/opentracing-go
RUN go get github.com/openzipkin/zipkin-go-opentracing
RUN go install github.com/abderraoufbensalem/cloudinary-service

ENTRYPOINT /go/bin/cloudinary-service

EXPOSE 8090
