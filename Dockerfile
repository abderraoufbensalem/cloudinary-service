FROM golang:latest
COPY ${TRAVIS_BUILD_DIR} ./bin/cloudinary-service
WORKDIR /go/bin/cloudinary-service
RUN ls -l
ENTRYPOINT /go/bin/cloudinary-service/cloudinary-service

EXPOSE 8080
