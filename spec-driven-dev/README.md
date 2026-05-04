# AI-learning

Spec-driven development (SDD) template for Go projects, usable with **Cursor** or **Claude Code** (and other AI coding agents).

## Use as a library (Go, Java, JS)

Same API and [file layout](docs/sdd/LIBRARY.md) in all bindings: path helpers + load spec/plan/tasks/progress from a base dir (e.g. `docs/sdd`).

### Go

```bash
go get github.com/buiviethoang/AI-learning/sdd
```

```go
import "github.com/buiviethoang/AI-learning/sdd"

spec, err := sdd.LoadSpec(ctx, "docs/sdd", "notification-service")
// sdd.SpecPath, sdd.PlanPath, sdd.TasksPath, sdd.ProgressPath, sdd.Paths(baseDir)
```

### Java (Maven)

```xml
<dependency>
  <groupId>ai.learning</groupId>
  <artifactId>sdd</artifactId>
  <version>0.1.0</version>
</dependency>
```

Install from repo: `mvn install` in `sdd-java/`, then depend on `ai.learning:sdd`.

```java
import ai.learning.sdd.SDD;
import ai.learning.sdd.Spec;

Spec spec = SDD.loadSpec("docs/sdd", "notification-service");
// SDD.specPath, planPath, tasksPath, progressPath, paths(baseDir)
```

### JS / Node

```bash
npm install /path/to/AI-learning/sdd-js
# or link: npm link ./sdd-js
```

```js
const sdd = require('sdd-sdk');
const spec = sdd.loadSpec('docs/sdd', 'notification-service');
// sdd.loadSpecAsync, loadPlan, loadTasks, loadProgress; path helpers: specPath, planPath, etc.
```

TypeScript: types in `sdd-js/index.d.ts`.

**Contract**: [docs/sdd/LIBRARY.md](docs/sdd/LIBRARY.md) — layout and API surface shared by all bindings.

## Quick links

- **SDD flow**: [docs/sdd/FLOW.md](docs/sdd/FLOW.md) — Specify → Plan → Tasks → Implement
- **Library contract**: [docs/sdd/LIBRARY.md](docs/sdd/LIBRARY.md) — layout and API for Go/Java/JS
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
