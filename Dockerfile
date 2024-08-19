FROM golang:1.23 AS builder
WORKDIR /app

COPY . .
RUN make install-deps

RUN CGO_ENABLED=0 GOOS=linux make build

EXPOSE 3000
CMD ["/app/dist/gooferr"]