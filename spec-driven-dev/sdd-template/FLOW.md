# Spec-Driven Development — Basic Flow

Four phases: **Specify → Plan → Tasks → Implement**. Use for production features, multi-file work, and multi-agent (e.g. Cursor) sessions.

## 1. Specify

- **Where**: `docs/sdd/specs/<feature>.md`
- **What**: User story, stakeholders, success criteria, functional/non-functional requirements, **explicit constraints (DO NOT)**, technical context, acceptance tests.
- **Rule**: Another developer (or agent) could implement without clarification.

**Prompt idea**: "Create a comprehensive specification for [feature] that includes: user story and stakeholders, measurable success criteria, functional and non-functional requirements, explicit constraints (what NOT to build), technical context and integration points, acceptance tests. Be specific enough that another developer could implement without clarification."

## 2. Plan

- **Where**: `docs/sdd/plans/<feature>-plan.md`
- **What**: Technical architecture, DB schema, API design, integration points. Must respect existing codebase and constraints from the spec.
- **Input**: Spec + (optional) repo/codebase context.

**Prompt idea**: "Based on the specification in `docs/sdd/specs/<feature>.md` and our existing codebase, create a technical implementation plan: architecture, database schema, API design, integration points. Respect all DO NOTs and existing patterns."

## 3. Tasks

- **Where**: In the plan doc or `docs/sdd/plans/<feature>-tasks.md`
- **What**: Discrete, independently implementable, testable units (single-PR scope). Dependencies and order.

**Prompt idea**: "Break down the implementation into discrete tasks. Each task: independently implementable, testable in isolation, reviewable as a single PR. Include dependencies and order."

## 4. Implement

- **Per task**: Implement with spec + plan in context. Reference spec section and acceptance criteria.
- **Validation**: Acceptance criteria from spec become tests. Failures = deviation from spec.
- **State**: Use `docs/sdd/progress/<feature>-progress.md` (and optionally test-status) so the next session or agent can resume.

**Resume prompt**: "Review `docs/sdd/progress/<feature>-progress.md` and test status. Continue the current task; keep consistency with the spec and plan."

---

## When to Use SDD vs Ad-Hoc Prompting

| Use ad-hoc prompting for | Use SDD for |
|--------------------------|-------------|
| Quick prototypes, experiments | Production features |
| Simple utilities, bug fixes | Multi-file / multi-session work |
| Learning, one-off scripts | Integration with existing systems |
| Small UI tweaks | Compliance or strict requirements |
| | Team or multi-agent consistency |

---

## File Layout (Template)

```
docs/sdd/
├── FLOW.md                    # This file
├── CLAUDE_CODE.md             # How to use this flow with Claude Code
├── specs/
│   ├── _template.md           # Minimal spec template
│   └── notification-service.md
├── plans/
│   ├── notification-service-plan.md
│   └── notification-service-tasks.md
└── progress/
    └── notification-service-progress.md

.cursor/
├── rules/                     # Cursor rules (SDD + Go)
└── skills/
    └── spec-driven-development/
        └── SKILL.md
```

---

## Verification

- Each acceptance criterion in the spec → test case.
- Flag deviations as `SPEC_DEVIATION: [reason]`.
- Before merge: implementation matches spec, DO NOTs respected, acceptance tests exist.
