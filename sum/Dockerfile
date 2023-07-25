FROM mcr.microsoft.com/oss/go/microsoft/golang:1.20-fips-cbl-mariner2.0 AS builder
WORKDIR /source

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o /app/sum .

# ------------------------------------------------------------------------------

# https://mcr.microsoft.com/en-us/product/cbl-mariner/distroless/base/about
# FROM mcr.microsoft.com/cbl-mariner/distroless/minimal:2.0 # Does not work, no glibc.

FROM mcr.microsoft.com/cbl-mariner/distroless/base:2.0
WORKDIR /app
COPY --from=builder /app .
ENTRYPOINT ["./sum"]
