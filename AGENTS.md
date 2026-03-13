# Multi-Agent / Cursor Workflow

This project uses **Spec-Driven Development (SDD)** so multiple agents or sessions can work consistently from one source of truth.

## For agents (Cursor, etc.)

1. **Spec first**  
   For features, read or create a spec under `docs/sdd/specs/`. See `docs/sdd/FLOW.md`.

2. **Rules**  
   - `.cursor/rules/sdd-standards.mdc` — always: check spec, implement against spec, flag deviations  
   - `.cursor/rules/go-standards.mdc` — when editing `**/*.go`: interfaces, table-driven tests, clean code  

3. **Skill**  
   - `.cursor/skills/spec-driven-development/` — when user asks for a feature or mentions specs/plans/tasks  

4. **Resuming**  
   - Read `docs/sdd/progress/<feature>-progress.md` and continue the current task.  

## Layout

```
docs/sdd/
  FLOW.md           # SDD flow (Specify → Plan → Tasks → Implement)
  specs/            # Feature specs (source of truth)
  plans/            # Technical plans and task lists
  progress/         # Per-feature progress for resuming

.cursor/
  rules/            # SDD + Go rules
  skills/           # spec-driven-development skill
```

## Example: notification service

- Spec: `docs/sdd/specs/notification-service.md`
- Plan: `docs/sdd/plans/notification-service-plan.md`
- Tasks: `docs/sdd/plans/notification-service-tasks.md`
- Progress: `docs/sdd/progress/notification-service-progress.md`

Use this as a template: copy the flow and folder structure for new features.

## Claude Code (no Cursor rules)

Claude Code doesn’t load `.cursor/rules` or skills. Use **`docs/sdd/CLAUDE_CODE.md`** for session setup, what to attach, and how to reuse the same rules as instructions.
