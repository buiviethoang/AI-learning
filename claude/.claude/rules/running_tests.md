# Running Tests

## Commands

```bash
make test                     # Unit tests only (fast)
make test-integration         # Integration tests (requires live DB/Kafka/Redis)
make test-all                 # Both
make cover                    # Coverage report (stdout)
make cover-html               # HTML coverage report → coverage.html
make cover-integration        # Coverage for integration tests
```

## Test Conventions

- Unit tests: standard `*_test.go`, no build tag, no external deps.
- Integration tests: require `//go:build integration` at top of file and a live environment.
- CI (`make ci`) runs only unit tests. Integration tests run separately.
- `cmd/worker` package tests are excluded from CI (legacy Machinery workers).

## Writing Tests

- Table-driven tests preferred (`tt.name` subtests via `t.Run`).
- Do not mock the DB in integration tests — use real Postgres.
- Seed data helpers are in `routes/*_seed_integration_test.go`.
- Test DB config is read from `config/config.local.yml` or env vars.
