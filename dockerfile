FROM golang:1.9.1 as builder
WORKDIR /go/src/github.com/martin-magakian/augmentor
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o augmentor .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/martin-magakian/augmentor/augmentor .
EXPOSE 3000
CMD ["./augmentor"]