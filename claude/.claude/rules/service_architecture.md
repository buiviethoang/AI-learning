# Service Architecture

## Binaries

| Binary | Role | Port |
|---|---|---|
| `receiver` | HTTP ingress — validates, resolves template, enqueues execution plan | 8080 |
| `ops` | Admin CRUD (tenants, markets, providers, templates, rules) | 8086 |
| `plan-scheduler` | Polls DB for due execution plans → dispatches to Kafka | — |
| `attempt-worker <channel>` | Per-channel Kafka consumer → calls provider SDK | — |
| `attempt-retry-forwarder` | Consumes retry topics with exponential backoff | — |
| `callback-providers` | Ingests provider webhooks, updates delivery status | — |
| `metrics-rollup` | Aggregates daily cost/delivery analytics | — |

## Request Lifecycle

```
Client → receiver (HTTP)
  → normalize → idempotency check → template resolve → routing profile
  → execution plan saved to DB
  → plan-scheduler polls → dispatches to Kafka attempt topic
  → attempt-worker consumes → provider send → outcome recorded
  → provider callback → callback-providers → delivery confirmed
```

## Routing & Provider Scoring

`internal/routing/` scores available providers per channel using:
- `provider_scoring_rules` (weight/priority config)
- `provider_success_window` (rolling success rate)
- `provider_success_daily` (daily stats)
- Allowlist rules and rate-limit rules

## Key Packages

- `internal/receiver/` — ingestion pipeline (normalize, idempotency, routing profile, template mgmt, engine)
- `internal/dispatch/` — execution plan state machine + Kafka dispatch
- `internal/worker/` — attempt runner, provider adapter, dedup, outcome handler
- `internal/callback/` — webhook ingestion, SMS delivery report, classify
- `internal/policy/` — quota enforcement, risk/blacklist control
- `internal/observability/` — lifecycle event emission (Kafka + Prometheus)

## External Dependencies

- **PostgreSQL + GORM** — primary store
- **Redis** — L1/L2 cache + legacy Machinery queue
- **Kafka (IBM/sarama)** — attempt topics, retry, DLQ, lifecycle events
- **Gin** — HTTP framework
- **OpenTelemetry + Elastic APM** — distributed tracing
- **Unleash** — feature flags
