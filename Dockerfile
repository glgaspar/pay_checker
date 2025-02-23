FROM golang:1.22.3
WORKDIR /src

ENV project_name=value

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /pay_checker

CMD ["/pay_checker"]