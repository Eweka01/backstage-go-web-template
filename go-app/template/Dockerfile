# --- build stage ---
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY src/go.mod ./
COPY src/main.go ./
COPY src/web ./web
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/app ./...

# --- runtime stage (~7-15 MB final image) ---
FROM scratch
COPY --from=build /out/app /app
EXPOSE 5000
ENTRYPOINT ["/app"]
