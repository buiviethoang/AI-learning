# Building the Project

## Prerequisites

- Go 1.25+
- Docker + Docker Compose (for local infra)
- `yq` (required for migrate targets)
- `golangci-lint`, `goimports`, `goose`, `swaggo/swag`

## Local Infra

```bash
docker-compose up -d          # PostgreSQL, Redis, Kafka, Zookeeper
docker-compose -f docker-compose.observability.yml up -d  # Jaeger, Prometheus, Grafana
```

## Config

Primary: `config/config.yml`. Local overrides: `config/config.local.yml` (gitignored).  
Env vars override YAML using `__` as separator: `POSTGRES__HOST=localhost`.

## Build & Run

```bash
make run-receiver             # HTTP ingress (port 8080)
make run-ops                  # Admin API (port 8086)
make run-worker-sms           # SMS attempt worker
make run-worker-email         # Email attempt worker
make run-worker-ott           # OTT worker (Zalo/WhatsApp)
make run-noti                 # Push notification worker
make run-metrics-rollup       # Daily analytics rollup
```

Each binary shares the same `cmd/main.go` entrypoint dispatched via cobra subcommands in `cmd/cmds.go`.

## Swagger

```bash
make swagger-receiver         # Regenerate docs/receiver/
make swagger-ops              # Regenerate docs/ops/
make swagger-callback         # Regenerate docs/callback/
```

## Git Hooks

```bash
./scripts/setup-githooks.sh   # Run once after clone
```

Hooks enforce on pre-commit: `gofmt`, `goimports`, `gitleaks`.  
On pre-push: `go test ./...`, `golangci-lint run`.
