FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o temperatureApp .

FROM scratch

COPY --from=builder /app/temperatureApp /app/temperatureApp

WORKDIR /app

CMD [ "./temperatureApp"]
