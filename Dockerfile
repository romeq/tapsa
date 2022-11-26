FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 8080

RUN apt update && apt install -y postgresql postgresql-client-14
RUN make setup build

CMD ["make", "migrateup", "run"]
