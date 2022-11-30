FROM --platform=$BUILDPLATFORM golang:1.19-buster AS build

WORKDIR /build

COPY go.mod go.sum Makefile ./

RUN go mod download

COPY . .

RUN make

FROM gcr.io/distroless/base-debian11

ARG TARGETPLATFORM

LABEL org.opencontainers.image.source=https://github.com/abhisek/sane
LABEL org.opencontainers.image.description="Git repository structure validator"
LABEL org.opencontainers.image.licenses=MIT

COPY --from=build /build/target/${TARGETPLATFORM}/sane /usr/local/bin/sane

USER nonroot:nonroot

ENTRYPOINT ["sane", "-v"]
