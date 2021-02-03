FROM golang:1.15-alpine as build

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
##Fix this part to COPY . .  if there is multiple subfolders
COPY /app .  

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

FROM alpine:latest 

COPY --from=build /dist/main /

ENTRYPOINT [ "/main" ]

EXPOSE 8080