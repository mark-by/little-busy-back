FROM golang:1.19.1 AS build
ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOFLAGS -mod=vendor
ARG BUILD_REF
ARG IMAGE
RUN mkdir /build
ADD . /build/

WORKDIR /build
RUN go build -o main_${IMAGE} ./${IMAGE}/cmd/main.go

FROM alpine:3.14
ARG BUILD_DATE
ARG BUILD_REF
ARG IMAGE
COPY --from=build /build/main_${IMAGE} /usr/bin
WORKDIR /${IMAGE}
ENTRYPOINT ["main_auth"]