# Order Notification Service Specification

**Context**: Go service (library or microservice). E-commerce order status notifications.

## User Story

As a customer, I want to receive timely updates about my order status through my preferred channels (email, SMS, in-app), so that I stay informed without being overwhelmed.

## Stakeholders

- **Primary**: E-commerce customers (50K daily active users)
- **Secondary**: Customer service (reduce "where's my order" tickets by 40%)
- **Tertiary**: Operations (monitor delivery performance)

## Success Criteria

1. **Delivery speed**: 95% of notifications delivered within 60 seconds
2. **Reliability**: 99.9% successful delivery rate after retries
3. **Preference compliance**: 100% adherence to user channel preferences
4. **Support impact**: 40% reduction in order-status tickets

## Functional Requirements

### Multi-Channel Support

- Email via SendGrid (or configurable provider)
- SMS via Twilio (or configurable provider)
- In-app via WebSocket or server-sent events
- Each channel independently toggleable

### Notification Types

1. Order confirmed  
2. Payment processed  
3. Order shipped (with tracking)  
4. Out for delivery  
5. Delivered  
6. Delivery failed  
7. Refund initiated  
8. Refund completed  

### User Preferences

- Global on/off
- Per-channel toggles
- Per-notification-type preferences
- Quiet hours: no SMS between 22:00–08:00 (user timezone)
- Language: English, Spanish, French

### Retry Logic

- 3 retry attempts for failed deliveries
- Exponential backoff: 1 min, 5 min, 15 min
- Strategy can differ per channel
- Dead letter queue after final failure

### Analytics

- Delivery rate per channel
- Open/click rates for email (if provider supports)
- Preference changes
- Retry patterns; alert on degradation

## Non-Functional Requirements

### Performance

- Handle 10,000 concurrent notifications
- 60s end-to-end delivery SLA
- API response time ≤ 100ms (enqueue/trigger only)

### Security

- PII encryption at rest
- Secure tokens for unsubscribe links
- Rate limit: max 50 notifications per user per day

### Compliance

- CAN-SPAM: one-click unsubscribe in email
- TCPA compliance for SMS
- GDPR-compliant data handling

## Explicit Constraints (DO NOT)

- Do NOT implement push (mobile push) in Phase 1
- Do NOT build a custom SMTP/email sender; use SendGrid (or injected provider)
- Do NOT modify the existing User/Identity model; integrate via interfaces
- Do NOT add social/messaging (e.g. WhatsApp, Facebook) in Phase 1
- Do NOT build admin UI for editing templates; admins use code/config
- Do NOT assume a separate deployable service; may be embedded in monolith or run as a service

## Technical Context (Go)

- Prefer interfaces for providers (EmailSender, SMSSender, InAppSender)
- Use `context.Context` for cancellation and timeouts
- Table-driven tests for business logic
- No global mutable state; dependencies injected
- Existing stack: assume PostgreSQL (or existing store), Redis for queue/cache if needed

## Acceptance Tests

1. User can enable/disable individual channels; only enabled channels receive.
2. Notification is enqueued and delivered within 60s of trigger (under load).
3. Failed sends retry with backoff 1 min, 5 min, 15 min; then DLQ.
4. Quiet hours: SMS not sent 22:00–08:00; queued for next window.
5. Unsubscribe link works without auth and persists preference.
6. Analytics expose delivery rate and retry metrics (e.g. metrics endpoint or callback).

## Test Scenarios

### Happy path

Given user has email and SMS enabled  
When order status changes to "shipped"  
Then email and SMS sent within 60s  

### Edge: Quiet hours

Given current time 23:00, user has SMS enabled  
When "delivered" notification triggered  
Then email and in-app sent immediately; SMS queued for 08:00  

### Error: Provider failure

Given SendGrid returns 500  
When email send attempted  
Then retry after 1 min, then 5 min, then 15 min; then DLQ  
