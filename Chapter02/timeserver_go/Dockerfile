FROM golang:latest as builder
WORKDIR /app
RUN go mod init timeserver
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /timeserver

FROM gcr.io/distroless/base-debian12
#FROM alpine:latest
WORKDIR /
COPY --from=builder /timeserver /timeserver
ENV PORT 80
#USER nonroot:nonroot
CMD ["/timeserver"]