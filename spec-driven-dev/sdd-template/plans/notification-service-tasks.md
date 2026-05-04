# Notification Service — Task Breakdown

**Plan**: [notification-service-plan.md](notification-service-plan.md)  
**Spec**: [notification-service.md](../specs/notification-service.md)

Each task: one PR, independently testable, reviewable against spec section.

| # | Task | Spec section | Deps |
|---|------|--------------|------|
| 1 | DB/schema for preferences and DLQ (or adapter to existing store) | Technical context, Retry | — |
| 2 | Core models (Notification, Preference, DeliveryResult); interfaces for providers | Technical context | 1 |
| 3 | Preference manager: load, apply, quiet hours, per-channel/per-type | User preferences, Quiet hours | 2 |
| 4 | Email provider (SendGrid); retry + analytics hooks | Multi-channel, Retry, Analytics | 2 |
| 5 | SMS provider (Twilio); quiet hours + retry | Multi-channel, Retry, Quiet hours | 2, 3 |
| 6 | In-app provider (WebSocket/SSE stub or real) | Multi-channel | 2 |
| 7 | Queue + worker: consume, call providers, retry backoff, DLQ | Retry logic, Performance | 2, 4, 5, 6 |
| 8 | Service orchestration: Notify(ctx, userID, eventType, payload) | All | 3, 7 |
| 9 | Unsubscribe flow (token, persistence) | Compliance, Acceptance 5 | 3 |
| 10 | Analytics and metrics (delivery rate, retries) | Analytics, Acceptance 6 | 7 |
| 11 | Table-driven tests for preferences, retry, quiet hours | Acceptance tests | 3, 4, 5, 7 |

Dependencies: implement in order 1 → 2 → 3; 4,5,6 in parallel after 2; 7 after 4,5,6; 8 after 7; 9,10 can follow 8.
