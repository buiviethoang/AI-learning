package sdd

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestPaths(t *testing.T) {
	tests := []struct {
		name       string
		base       string
		wantSpecs  string
		wantPlans  string
		wantProgress string
	}{
		{"relative", "docs/sdd", "docs/sdd/specs", "docs/sdd/plans", "docs/sdd/progress"},
		{"absolute", "/repo/docs/sdd", "/repo/docs/sdd/specs", "/repo/docs/sdd/plans", "/repo/docs/sdd/progress"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			specs, plans, progress := Paths(tt.base)
			if specs != tt.wantSpecs || plans != tt.wantPlans || progress != tt.wantProgress {
				t.Errorf("Paths() = %q, %q, %q; want %q, %q, %q", specs, plans, progress, tt.wantSpecs, tt.wantPlans, tt.wantProgress)
			}
		})
	}
}

func TestSpecPath(t *testing.T) {
	tests := []struct {
		base    string
		feature string
		want    string
	}{
		{"docs/sdd", "notification-service", filepath.Join("docs", "sdd", "specs", "notification-service.md")},
		{".", "foo", filepath.Join("specs", "foo.md")},
	}
	for _, tt := range tests {
		got := SpecPath(tt.base, tt.feature)
		if got != tt.want {
			t.Errorf("SpecPath(%q, %q) = %q, want %q", tt.base, tt.feature, got, tt.want)
		}
	}
}

func TestProgressPath(t *testing.T) {
	got := ProgressPath("docs/sdd", "notification-service")
	want := filepath.Join("docs", "sdd", "progress", "notification-service-progress.md")
	if got != want {
		t.Errorf("ProgressPath() = %q, want %q", got, want)
	}
}

func TestLoadSpec(t *testing.T) {
	dir := t.TempDir()
	specsDir := filepath.Join(dir, SpecsDirName)
	_ = os.MkdirAll(specsDir, 0755)
	path := filepath.Join(specsDir, "foo.md")
	content := []byte("# Foo Spec\n\n## User Story\n\nAs a user...")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		baseDir string
		feature string
		want    *Spec
		wantErr bool
	}{
		{"ok", dir, "foo", &Spec{Feature: "foo", Content: content}, false},
		{"missing", dir, "nonexistent", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadSpec(context.Background(), tt.baseDir, tt.feature)
			if (err != nil) != tt.wantErr {
				t.Fatalf("LoadSpec() err = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got.Feature != tt.want.Feature || !reflect.DeepEqual(got.Content, tt.want.Content) {
				t.Errorf("LoadSpec() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestLoadProgress(t *testing.T) {
	dir := t.TempDir()
	progressDir := filepath.Join(dir, ProgressDirName)
	_ = os.MkdirAll(progressDir, 0755)
	path := filepath.Join(progressDir, "foo-progress.md")
	content := []byte("# Progress\n\n## Completed\n- Task 1")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}

	got, err := LoadProgress(context.Background(), dir, "foo")
	if err != nil {
		t.Fatalf("LoadProgress() err = %v", err)
	}
	if got.Feature != "foo" || !reflect.DeepEqual(got.Content, content) {
		t.Errorf("LoadProgress() = %+v, content %q", got, got.Content)
	}
}
