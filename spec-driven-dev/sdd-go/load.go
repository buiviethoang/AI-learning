package sdd

import (
	"context"
	"os"
)

// LoadSpec reads the spec for feature from baseDir (e.g. docs/sdd).
func LoadSpec(ctx context.Context, baseDir, feature string) (*Spec, error) {
	path := SpecPath(baseDir, feature)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Spec{Feature: feature, Content: content}, nil
}

// LoadPlan reads the plan for feature from baseDir.
func LoadPlan(ctx context.Context, baseDir, feature string) (*Plan, error) {
	path := PlanPath(baseDir, feature)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Plan{Feature: feature, Content: content}, nil
}

// LoadTasks reads the tasks file for feature from baseDir.
func LoadTasks(ctx context.Context, baseDir, feature string) ([]byte, error) {
	path := TasksPath(baseDir, feature)
	return os.ReadFile(path)
}

// LoadProgress reads the progress file for feature from baseDir.
func LoadProgress(ctx context.Context, baseDir, feature string) (*Progress, error) {
	path := ProgressPath(baseDir, feature)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Progress{Feature: feature, Content: content}, nil
}
