FROM golang:1.22.2-bookworm

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o tfl .


EXPOSE 6060

CMD [ "./tfl" ]
