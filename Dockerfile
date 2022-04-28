### FIRST STAGE BUILD
FROM golang:alpine AS build

# install some awesome packages
RUN apk update && apk add git

# change the current working directory
WORKDIR /go/src/gitlab.com/bravilogy/red-badger

# install project dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy the application
COPY . .

# build the project
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .



### SECOND STAGE OF THE BUILD
FROM alpine:latest

# change the current working directory
WORKDIR /usr/local/bin

# copy the release file from the build stage
COPY --from=build /go/src/gitlab.com/bravilogy/red-badger/app .

CMD ["/usr/local/bin/app"]