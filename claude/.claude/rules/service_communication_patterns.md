# Service Communication Patterns

## Kafka Topics (attempt flow)

```
attempt.<channel>          Main attempt topic (receiver → attempt-worker)
attempt.<channel>.retry    Retry topic with exponential backoff delay
attempt.<channel>.dlq      Dead-letter queue after max retries exhausted
lifecycle.events           Observability events (delivery status changes)
```

`attempt-retry-forwarder` consumes retry topics, applies backoff, and re-publishes to the main topic.

## Kafka Consumer Pattern

Each `attempt-worker` binary is a Kafka consumer group per channel. Workers:
1. Consume message from topic
2. Check `worker_attempt_idempotency` (Redis + DB) — skip if already processed
3. Resolve provider credentials from config/DB
4. Call provider SDK (with timeout)
5. Record outcome in `notification_attempts`
6. Emit lifecycle event
7. On failure: publish to retry/DLQ topic

## HTTP — Internal Service Calls

Clients in `httpclient/` wrap outbound HTTP calls to:
- `account_service` — user profile lookup
- `media_service` — media URL resolution
- `translation_global_service` — i18n content
- `super_app_service` — super-app platform integration
- `receiver_service` — internal replay/re-send

All clients use the shared `httpclient.http_client.go` base with timeout and tracing headers.

## Callback / Webhook Ingestion

`callback-providers` exposes HTTP endpoints for provider delivery receipts.  
Routes defined in `routes/callback_attempt.go`.  
Each provider's callback is classified (`internal/callback/classify/`) and mapped to an attempt outcome.

## Caching Strategy

- **L1**: in-process `sync.Map` with TTL (routing profiles, provider config)
- **L2**: Redis (shared across replicas)
- Cache-aside pattern: check L1 → L2 → DB, write-through on miss.
- Redis also used for rate-limit counters (atomic INCR with TTL).
