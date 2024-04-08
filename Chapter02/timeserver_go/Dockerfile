FROM golang:1.22.0 as builder
WORKDIR /app
RUN go mod init timeserver
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /timeserver

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=builder /timeserver /timeserver
ENV PORT 8080
USER nonroot:nonroot
CMD ["/timeserver"]