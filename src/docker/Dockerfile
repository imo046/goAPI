FROM golang:1.19-alpine3.15 AS builder

ENV GO111MODULE=on

WORKDIR /home

COPY . .

#dependencies
RUN go mod tidy

WORKDIR /home/cmd/tutorial

RUN GOOS=linux go build -a -installsuffix cgo -o app .


#Second stage for a smaller image
FROM alpine:3.15
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /home/cmd/tutorial/app .
CMD ["./app"]