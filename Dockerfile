# Build stage
FROM golang AS build-env
WORKDIR /go/src/akovalyov/chlorine/
ADD . ./

# Install project dependencies
RUN go get golang.org/x/oauth2/clientcredentials
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/sessions
RUN go get github.com/zmb3/spotify

# Compile application
RUN CGO_ENABLED=0 go build -o app

# Final stage
FROM alpine AS finalize

EXPOSE 8080

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-env /go/src/akovalyov/chlorine/app .
ENTRYPOINT ./app