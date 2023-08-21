FROM golang:latest AS BUILD

WORKDIR /build

COPY go.mod .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /broempSignal_Bots cmd/broempSignal_Bots.go

FROM alpine
WORKDIR /app
COPY --from=BUILD /broempSignal_Bots .
CMD ["/app/broempSignal_Bots"]
