# Service Architecture

## Binaries

| Binary | Role | Port |
|---|---|---|
| `receiver` | HTTP ingress — validates, idempotency check, risk policy → publishes `IngestMessage` to Kafka → HTTP 202 | 8080 |
| `ops` | Admin CRUD (tenants, markets, providers, templates, rules) | 8086 |
| `request-ingestor` | Kafka consumer `notif.request.ingest.v1`; writes `notification_requests` + `execution_plans`; calls Routing Engine + Dispatcher | — |
| `attempt-worker <channel>` | Per-channel Kafka consumer → calls provider SDK | — |
| `attempt-retry-forwarder` | Consumes retry topics with exponential backoff | — |
| `callback-providers` | Ingests provider webhooks, updates delivery status | — |
| `attempt-audit-consumer` | Kafka consumer `notif.attempt.audit.v1` → batch-inserts `notification_attempts` | — |
| `request-audit-consumer` | Kafka consumer `notif.request.audit.v1` → batch-inserts `notification_requests` | — |
| `plan-terminal-consumer` | Kafka consumer `notif.plan.terminal.v1` → updates `execution_plans` terminal state | — |
| `metrics-rollup` | Aggregates daily cost/delivery analytics | — |

## Request Lifecycle

```
Client → receiver (HTTP)
  → normalize → idempotency check → template resolve → routing profile
  → execution plan saved to DB
  → request-ingestor calls Dispatcher → publishes to Kafka attempt topic
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
