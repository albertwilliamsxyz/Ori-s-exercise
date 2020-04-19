FROM golang:latest as builder
LABEL maintainer="Albert Williams <albertwilliams.xyz@gmail.com>"
WORKDIR /technical-exercise
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/main ./cmd/server/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /technical-exercise/bin/main .
EXPOSE 50051
ENV TECHNICAL_EXERCISE_SERVER_ADDRESS=":50051"
CMD ["./main"] 
