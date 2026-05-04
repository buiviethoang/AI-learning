# Using This Repo with Claude Code

Claude Code does not read `.cursor/rules` or `.cursor/skills`. Use this guide to get the same spec-driven development workflow by supplying context and instructions yourself.

---

## 1. Session setup

At the start of a session, tell Claude:

> This repo uses spec-driven development. Read `docs/sdd/FLOW.md` and `AGENTS.md`. For features we: Specify → Plan → Tasks → Implement. Specs are in `docs/sdd/specs/`, plans in `docs/sdd/plans/`, progress in `docs/sdd/progress/`.

Optionally attach or paste:

- `docs/sdd/FLOW.md`
- `AGENTS.md`

---

## 2. Starting a new feature

**Option A — Ask Claude to write the spec**

> We need to add [feature]. Follow spec-driven development: create a spec from the template in `docs/sdd/specs/_template.md`, then a technical plan and a task list. Put them in `docs/sdd/specs/`, `docs/sdd/plans/`, and optionally `docs/sdd/progress/`.

**Option B — You already have a spec**

> Implement from the spec at `docs/sdd/specs/<feature>.md`. Create a plan and task list in `docs/sdd/plans/` if they don’t exist, then implement the first task.

---

## 3. Implementing from an existing spec

Attach or paste:

- `docs/sdd/specs/notification-service.md` (or your feature spec)
- `docs/sdd/plans/notification-service-plan.md`
- `docs/sdd/plans/notification-service-tasks.md`

Then say:

> Implement task [N] from the task list. Follow the spec strictly (requirements, DO NOTs, acceptance criteria). Use Go: interfaces for providers, `context.Context`, table-driven tests. See `.cursor/rules/go-standards.mdc` for conventions.

To mimic the SDD rule, add:

> Before coding, confirm the task matches the spec. Turn acceptance criteria into tests. If you must deviate from the spec, add a comment: `// SPEC_DEVIATION: reason`.

---

## 4. Resuming work

Attach:

- `docs/sdd/progress/notification-service-progress.md` (or your feature’s progress file)
- The spec and task list (if not already in context)

Then say:

> Continue from this progress file. The current task is [name/number]. Stay consistent with the spec and plan. Update the progress file when done.

---

## 5. Reusing Cursor rules as instructions

To get behavior similar to the Cursor rules in Claude Code, attach or paste these and ask Claude to follow them for the session:

| File | Purpose |
|------|---------|
| `docs/sdd/FLOW.md` | SDD phases and file layout |
| `.cursor/rules/sdd-standards.mdc` | Check spec first, implement against spec, SPEC_DEVIATION |
| `.cursor/rules/go-standards.mdc` | Go style: interfaces, table-driven tests, context, errors |

Prompt:

> Follow the rules in these files for this session: check for a spec before implementing a feature, implement against the spec, use table-driven tests in Go, and add SPEC_DEVIATION if you must diverge.

---

## 6. Quick reference — files to attach by scenario

| Scenario | Attach / reference |
|----------|---------------------|
| New feature (full SDD) | `FLOW.md`, `specs/_template.md` |
| Implement one task | Spec, plan, task list, (optional) `go-standards.mdc` |
| Resume | Progress file, spec, task list |
| “Same as Cursor rules” | `FLOW.md`, `sdd-standards.mdc`, `go-standards.mdc` |

---

## 7. Example: notification service task 4

1. Attach: `docs/sdd/specs/notification-service.md`, `docs/sdd/plans/notification-service-tasks.md`.
2. Say:

   > Implement task 4 (Email provider) from the task list. Follow the spec’s Retry Logic and Multi-Channel sections. Use Go: interface for the email sender, table-driven tests, context.Context. Add tests for the acceptance criteria that touch email delivery.

Claude Code will not auto-load rules, so the explicit instructions in this guide replace what Cursor does automatically.
