# SDD library contract

All language bindings (Go, Java, JS) use the same **file layout** and **API surface** so you can load specs from any project.

## File layout (under a base dir, e.g. `docs/sdd`)


| Path                                    | Description             |
| --------------------------------------- | ----------------------- |
| `{base}/specs/{feature}.md`             | Feature specification   |
| `{base}/plans/{feature}-plan.md`        | Technical plan          |
| `{base}/plans/{feature}-tasks.md`       | Task breakdown          |
| `{base}/progress/{feature}-progress.md` | Implementation progress |


Directory names: `specs`, `plans`, `progress`.

## API surface (same in every binding)

- **Path helpers** (no I/O): `specPath(baseDir, feature)`, `planPath(baseDir, feature)`, `tasksPath(baseDir, feature)`, `progressPath(baseDir, feature)`. Return the full path string.
- **Paths**: `paths(baseDir)` → `{ specsDir, plansDir, progressDir }`.
- **Loaders** (read from disk): `loadSpec`, `loadPlan`, `loadTasks`, `loadProgress`. Each takes `(baseDir, feature)` and returns content; loadSpec/loadPlan/loadProgress return `{ feature, content }`, loadTasks returns raw content.
- **Types**: `Spec`, `Plan`, `Progress` each have `feature: string` and `content: string` (or bytes). Tasks is raw content only.

## Bindings

- **Go**: `github.com/buiviethoang/AI-learning/sdd` — `LoadSpec(ctx, baseDir, feature)`, path helpers, `Spec.Content` is `[]byte`.
- **Java**: `sdd-java/` — Maven artifact; same path and load API; `Spec.getContent()` returns `String` or `byte[]`.
- **JS**: `sdd-js/` — npm-style package; sync `loadSpec(baseDir, feature)`; optional async where applicable; TypeScript types included.

