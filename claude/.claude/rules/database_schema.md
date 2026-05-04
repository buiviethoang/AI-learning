# Database Schema

Managed by **Goose** (`migrations/`). All tables are in the `public` schema. UUIDs via `uuid-ossp`.

## Core Tables

| Table | Purpose |
|---|---|
| `tenants` | Top-level org identity (`tenant_key` unique) |
| `markets` | Per-tenant market (`market_code`, country, timezone, currency) |
| `providers` | Provider registry (`provider_key`, `channel_group`, config schema) |
| `templates` | Versioned message templates per tenant/market/channel/locale |
| `template_routing_profile` | Maps template key → allowed channels + provider scoring profile |

## Notification Flow Tables

| Table | Purpose |
|---|---|
| `notification_requests` | Deduplication + idempotency record (one per external request) |
| `execution_plans` | Planned delivery attempt (channel, provider, scheduled_at, status) |
| `notification_attempts` | Individual provider call records (status, provider response, cost) |
| `worker_attempt_idempotency` | Prevents double-processing of Kafka messages |

## Policy / Scoring Tables

| Table | Purpose |
|---|---|
| `rate_limit_rules` | Per-tenant/channel rate caps |
| `blacklist_user_id_rules` | User IDs blocked from receiving notifications |
| `provider_scoring_rules` | Weight/priority rules for provider selection |
| `provider_success_daily` | Daily success/failure counters per provider |
| `provider_success_window` | Rolling-window success rate (used for real-time scoring) |

## Device Tokens

`device_tokens` is **hash-partitioned** into 8 shards (MODULUS 8, REMAINDER 0–7).  
Always include a partition key (`user_id` or `device_id`) in queries or you'll scan all 8 shards.

## Migrations

```bash
make migrate-status     # show applied/pending
make migrate-up         # apply all pending
make migrate-down       # rollback last
make migrate-new        # scaffold new migration file
```
