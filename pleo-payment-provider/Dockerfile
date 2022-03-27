FROM golang:1.16-buster AS build

WORKDIR /app

# COPY go.mod ./
# RUN go mod download

COPY *.go go.mod ./

RUN go build -o /output


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /output /output

EXPOSE 9000

USER nonroot:nonroot

ENTRYPOINT ["/output"]
