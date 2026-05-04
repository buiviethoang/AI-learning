# Code Conventions

## Language & Formatting

- Go 1.25. Run `make fmt` (`gofmt`) and `make verify` before committing.
- `golangci-lint` is the linter — config at `.golangci.yml`. Run `make lint`.
- `goimports` manages import grouping (stdlib / external / internal).

## Package Structure (per `internal/*` domain)

```
transport/http/    Gin handlers — parse request, call service, render response
business/service/  Domain logic — orchestration, no HTTP/DB knowledge
storage/repo/      DB interfaces and GORM implementations
dto/               Request/response structs (transport ↔ service boundary)
contract/          Shared types and repo/provider interfaces
model/             DB model structs (GORM tags)
```

## Dependency Injection

- Every service/handler constructor accepts a `*Deps` struct (not variadic args).
- Missing required deps → `panic` at startup (fail-fast).
- Interfaces live in `contract/` or alongside the consumer; implementations in `storage/` or `pkg/`.

## Error Handling

- Return `error` up the stack; wrap with `fmt.Errorf("context: %w", err)`.
- HTTP errors use structured error responses via the shared error helper in `internal/httpserver/`.
- Ops-specific errors use `internal/opsadmin/opserr`.

## Naming

- Exported types: `PascalCase`. Unexported: `camelCase`.
- Repo interfaces: `XxxRepository`. Service interfaces: `XxxService`.
- Kafka consumer structs: `XxxWorker` or `XxxConsumer`.

## Comments

- Only write comments when the WHY is non-obvious. No docstrings on trivial functions.

## Feature Flags

- Runtime toggles via Unleash, accessed through the shared feature-flag client in `pkg/`.
- New behaviour that is risky should be gated behind a flag.
