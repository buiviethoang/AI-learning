# CLAUDE.md

## Project

Multi-binary notification platform (Go) that routes and delivers messages across SMS, Email, OTT (Zalo/WhatsApp), Push (FCM/APNS/Huawei), Pushkit, and iOS Live Activity. Supports multi-tenant, multi-market deployments with intelligent provider scoring, retry/backoff, and delivery callbacks.

## Commands

```bash
# Run binaries
make run-receiver          # HTTP ingress (port 8080)
make run-ops               # Admin server (port 8086)
make run-worker-sms|email|ott|noti  # Channel workers

# Quality & CI
make fmt && make lint      # Format + lint
make ci                    # Full CI: verify + test + lint + security

# Database
make migrate-up|down|status|new

# Swagger
make swagger-receiver|ops|callback
```

For full build/test details → `./.claude/rules/building_the_project.md` and `running_tests.md`.

## Architecture

```
cmd/               → binary entrypoints
internal/          → V2 clean architecture (receiver, routing, dispatch, worker, callback, opsadmin, metrics, policy)
module/            → V1 legacy (Machinery task queue)
pkg/               → shared infra clients (postgres, redis, kafka, fcm, apns…)
config/            → Viper config (config.yml / config.local.yml)
migrations/        → Goose SQL migrations
```

Layer pattern per `internal/*` package: `transport/http → business/service → storage/repo`, with `dto/` and `contract/` at each layer boundary. Constructors take `*Deps` and panic on missing deps.

For service flow, Kafka topics, and provider integration → `./.claude/rules/service_architecture.md` and `service_communication_patterns.md`.  
For DB schema → `./.claude/rules/database_schema.md`.  
For coding conventions → `./.claude/rules/code_conventions.md`.

## Important Notes

- Run `./scripts/setup-githooks.sh` once after cloning (enforces gofmt, goimports, gitleaks, tests, lint).
- Config overrides via env: `POSTGRES__HOST=localhost` (double underscore for nested keys).
- Integration tests require `//go:build integration` tag; excluded from `make test`.
- `device_tokens` is hash-partitioned (8 shards) — never query without a partition key.
- `module/` (V1) is legacy; all new features go into `internal/` (V2).
