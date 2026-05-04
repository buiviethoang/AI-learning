# Notification Service — Technical Plan

**Spec**: [notification-service.md](../specs/notification-service.md)

## Architecture Overview

```
Triggers (Order/Event) → Notification Core → Preference Check → Queue
                                                                    ↓
Storage (Preferences, DLQ)    Providers: Email (SendGrid), SMS (Twilio), InApp (WS)
```

- **Notification core**: orchestration, preference resolution, channel selection
- **Queue**: in-memory or Redis; worker consumes and calls providers
- **Providers**: interfaces `EmailSender`, `SMSSender`, `InAppSender`; implementations call external APIs

## Go Package Layout (Suggested)

```
internal/notification/
  service.go       # orchestration
  preferences.go   # load/apply user preferences, quiet hours
  queue.go         # enqueue / worker (or use existing queue pkg)
  providers/
    interface.go   # Sender interfaces
    email.go       # SendGrid
    sms.go         # Twilio
    inapp.go       # WebSocket/SSE
  models.go        # Notification, Preference, DeliveryResult
  retry.go         # backoff + DLQ
```

## Database / Storage

- Preferences: table or key-value (user_id, channel, notification_types, quiet_hours, language)
- DLQ: table or queue (payload, failure_reason, attempts, next_retry_at)
- Analytics: counters/metrics (delivery_ok, delivery_fail, retries) — export to existing metrics or DB

## API Design

- **Trigger**: `Notify(ctx, userID, eventType, payload)` — enqueue; return quickly (<100ms)
- **Preferences**: `GetPreferences(ctx, userID)`, `UpdatePreferences(ctx, userID, prefs)`; optional REST if needed
- **Unsubscribe**: token in link → validate → set channel or global off

## Integration Points

- Order/event system: calls `Notify` on status change
- User service: resolve user contact info and timezone (via interface if possible)
- Existing auth: optional for unsubscribe (token-based preferred so no auth required)

## Constraints Checklist

- [ ] No push notifications (Phase 1)
- [ ] Use SendGrid (or injected provider), no custom SMTP
- [ ] No changes to existing User model; use interfaces
- [ ] No social channels (Phase 1)
- [ ] No template-editing UI; code/config only
- [ ] Works as library or service (no hard dependency on deploy topology)
