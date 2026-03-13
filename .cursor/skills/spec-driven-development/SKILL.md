---
name: spec-driven-development
description: Follow spec-driven development (SDD) for production features. Use when the user asks for a new feature, multi-file implementation, or work with specs, plans, or tasks in docs/sdd/. Apply for Cursor or multi-agent workflows.
---

# Spec-Driven Development

## When to use

- User asks to build a **feature** (e.g. notification service, API, integration)
- Work spans **multiple files** or **multiple sessions**
- Repo contains `docs/sdd/` with specs or user mentions "spec" or "SDD"

## Four phases

1. **Specify** — `docs/sdd/specs/<feature>.md`: user story, success criteria, requirements, DO NOTs, acceptance tests
2. **Plan** — `docs/sdd/plans/<feature>-plan.md`: architecture, packages, API, storage
3. **Tasks** — `docs/sdd/plans/<feature>-tasks.md`: discrete, testable, single-PR tasks with dependencies
4. **Implement** — per task; reference spec section and acceptance criteria; update `docs/sdd/progress/<feature>-progress.md`

## Before coding a feature

1. Look for existing spec in `docs/sdd/specs/`. If none, offer to create one from `docs/sdd/specs/_template.md`.
2. If a plan exists, use it for package layout and integration points.
3. If tasks exist, implement or continue the current task; keep progress file updated.

## During implementation

- Cite spec and section (e.g. "per notification-service.md § Retry Logic").
- Turn acceptance criteria into table-driven tests (Go).
- Mark deviations: `// SPEC_DEVIATION: reason`

## Resuming work

- Read `docs/sdd/progress/<feature>-progress.md` and any test-status file.
- Continue the current task; stay consistent with spec and plan.

## Reference

- Flow and file layout: `docs/sdd/FLOW.md`
- Example spec: `docs/sdd/specs/notification-service.md`
- Template: `docs/sdd/specs/_template.md`
