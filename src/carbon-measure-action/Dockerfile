FROM golang:1.17.7-alpine AS action-builder
WORKDIR /actioncompile
COPY go.mod ./
#COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . .
RUN go build -o carbon_measure .

FROM alpine:latest AS prod
WORKDIR /build
RUN mkdir -p /client/build
COPY --from=action-builder /actioncompile/carbon_measure .

# Declare entrypoints and activation commands.
ENTRYPOINT ["./carbon_measure"]