FROM golang:1.12 as builder
WORKDIR /app
COPY . .
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN make build

FROM scratch
WORKDIR /app
COPY --from=builder /app/bin/k3d-tools .
ENTRYPOINT [ "./k3d-tools" ]
