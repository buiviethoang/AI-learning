package sdd

import "path/filepath"

const (
	SpecsDirName    = "specs"
	PlansDirName    = "plans"
	ProgressDirName = "progress"
)

// Spec holds a feature specification (markdown content and metadata).
type Spec struct {
	Feature string
	Content []byte
}

// Plan holds a technical plan (markdown).
type Plan struct {
	Feature string
	Content []byte
}

// Progress holds implementation progress (markdown).
type Progress struct {
	Feature string
	Content []byte
}

// Paths returns standard SDD subdirs under base (e.g. docs/sdd).
func Paths(base string) (specs, plans, progress string) {
	return filepath.Join(base, SpecsDirName),
		filepath.Join(base, PlansDirName),
		filepath.Join(base, ProgressDirName)
}

// SpecPath returns the path to a spec file: base/specs/<feature>.md
func SpecPath(base, feature string) string {
	return filepath.Join(base, SpecsDirName, feature+".md")
}

// PlanPath returns the path to a plan file: base/plans/<feature>-plan.md
func PlanPath(base, feature string) string {
	return filepath.Join(base, PlansDirName, feature+"-plan.md")
}

// TasksPath returns the path to a tasks file: base/plans/<feature>-tasks.md
func TasksPath(base, feature string) string {
	return filepath.Join(base, PlansDirName, feature+"-tasks.md")
}

// ProgressPath returns the path to a progress file: base/progress/<feature>-progress.md
func ProgressPath(base, feature string) string {
	return filepath.Join(base, ProgressDirName, feature+"-progress.md")
}
