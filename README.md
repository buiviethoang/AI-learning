# AI-learning

Spec-driven development (SDD) template for Go projects, usable with **Cursor** or **Claude Code** (and other AI coding agents).

## Quick links

- **SDD flow**: [docs/sdd/FLOW.md](docs/sdd/FLOW.md) — Specify → Plan → Tasks → Implement
- **For agents (Cursor)**: [AGENTS.md](AGENTS.md) — rules, skills, resuming
- **Example spec**: [docs/sdd/specs/notification-service.md](docs/sdd/specs/notification-service.md)
- **Using Claude Code**: [docs/sdd/CLAUDE_CODE.md](docs/sdd/CLAUDE_CODE.md)

---

## How to apply

### In Cursor

Rules and the SDD skill are loaded automatically.

- **SDD rule** (`.cursor/rules/sdd-standards.mdc`) — `alwaysApply: true`; agent checks `docs/sdd/specs/` and implements against the spec.
- **Go rule** (`.cursor/rules/go-standards.mdc`) — applies when `*.go` files are in context.
- **Skill** (`.cursor/skills/spec-driven-development/`) — applies when you talk about features, specs, or multi-file work; or say “use spec-driven development”.

**In practice:**

| Goal | What to say |
|------|-------------|
| **New feature** | “We need to add [feature]. Follow spec-driven development: create a spec from `docs/sdd/specs/_template.md`, then a plan and tasks.” |
| **Implement from spec** | “Implement task 2 from `docs/sdd/plans/notification-service-tasks.md` per `docs/sdd/specs/notification-service.md`.” |
| **Resume work** | “Resume the notification service. Read `docs/sdd/progress/notification-service-progress.md` and continue the current task.” |
| **Keep agent on-spec** | “Implement the Email provider per notification-service.md, Retry Logic and Multi-Channel sections.” |

### In Claude Code

Claude Code does not use Cursor’s rules or skills; you supply context yourself. Use the dedicated guide:

- **[docs/sdd/CLAUDE_CODE.md](docs/sdd/CLAUDE_CODE.md)** — session setup, attaching specs/plans/progress, and reusing rules as instructions.

Summary: at session start, point Claude at `docs/sdd/FLOW.md` and `AGENTS.md`; attach the relevant spec (and plan/tasks/progress); ask it to follow the flow and implement. For full prompts and file list, see the guide.
