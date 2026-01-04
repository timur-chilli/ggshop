FROM golang:1.25.5-alpine AS base
WORKDIR /src


COPY go.mod go.sum ./
RUN go mod download


FROM base AS customer-side-builder
COPY customer_side ./customer_side
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /out/customer_side_service ./customer_side/cmd/app

FROM base AS warehouse-builder
COPY warehouse ./warehouse
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /out/warehouse_service ./warehouse/cmd/app

# --- runtime images ---
FROM alpine:3.20 AS customer-side
WORKDIR /app
COPY --from=customer-side-builder /out/customer_side_service /app/customer_side_service
ENTRYPOINT ["/app/customer_side_service"]

FROM alpine:3.20 AS warehouse
WORKDIR /app
COPY --from=warehouse-builder /out/warehouse_service /app/warehouse_service
ENTRYPOINT ["/app/warehouse_service"]