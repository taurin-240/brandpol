FROM golang:1.17-alpine

WORKDIR /src/

COPY go.mod go.sum ./
RUN go mod download

COPY ./src ./src/

# Starting app
CMD ["go", "run", "main.go settings.go database.go models.go pythonQuery.go"]