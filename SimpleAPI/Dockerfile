FROM golang:latest

LABEL maintainer="zieglernattacatalyst"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .



RUN go build 

# remove src files
RUN find . -name "*.go" -type f -delete

EXPOSE $PORT

CMD ["./SimpleAPI"]