# Compile stage
FROM golang:alpine as build

ADD . /gobuild
WORKDIR /gobuild
RUN go build -o jenkcli

# Application stage
FROM alpine

COPY --from=build /gobuild/jenkcli /jenkcli

ENTRYPOINT ["./jenkcli"]
CMD ["help"]