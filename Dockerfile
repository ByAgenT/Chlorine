# Build stage
FROM golang AS build-env
WORKDIR /go/src/chlorine/
ADD . ./

# Install project dependencies
RUN go get -d

# Compile application
RUN CGO_ENABLED=0 go build -o app

# Final stage
FROM alpine AS finalize

EXPOSE 8080

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-env /go/src/chlorine/app .
ENTRYPOINT ./app