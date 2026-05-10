package todo

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	var l TodoList
	l.Add("task 1")
	if len(l) != 1 {
		t.Fatalf("expected 1 item, got %d", len(l))
	}
	if l[0].Task != "task 1" {
		t.Fatalf("expected task 'task 1', got %q", l[0].Task)
	}
	if l[0].Done {
		t.Fatalf("new task should not be done")
	}
	if l[0].CreatedAt.IsZero() {
		t.Fatalf("CreatedAt should be set")
	}
	if time.Since(l[0].CreatedAt) > 2*time.Second {
		t.Fatalf("CreatedAt timestamp is not recent: %v", l[0].CreatedAt)
	}
	if !l[0].CompletedAt.IsZero() {
		t.Fatalf("CompletedAt should be zero for new tasks, got %v", l[0].CompletedAt)
	}
}

func TestComplete(t *testing.T) {
	var l TodoList
	l.Add("task 1")
	if err := l.Complete(1); err != nil {
		t.Fatalf("unexpected error completing item: %v", err)
	}
	if !l[0].Done {
		t.Fatalf("expected item to be marked done")
	}
	if l[0].CompletedAt.IsZero() {
		t.Fatalf("CompletedAt should be set after completing an item")
	}
	if time.Since(l[0].CompletedAt) > 2*time.Second {
		t.Fatalf("CompletedAt timestamp is not recent: %v", l[0].CompletedAt)
	}
	if l[0].CompletedAt.Before(l[0].CreatedAt) {
		t.Fatalf("CompletedAt should be after CreatedAt")
	}
	if err := l.Complete(2); err == nil {
		t.Fatalf("expected error when completing non-existent item")
	}
}

func TestDelete(t *testing.T) {
	var l TodoList
	l.Add("a")
	l.Add("b")
	if err := l.Delete(1); err != nil {
		t.Fatalf("unexpected error deleting item: %v", err)
	}
	if len(l) != 1 {
		t.Fatalf("expected 1 item after delete, got %d", len(l))
	}
	if l[0].Task != "b" {
		t.Fatalf("expected remaining task 'b', got %q", l[0].Task)
	}
	if err := l.Delete(0); err == nil {
		t.Fatalf("expected error when deleting with invalid index")
	}
}

func TestSaveAndGet(t *testing.T) {
	var l TodoList
	l.Add("persist")

	dir := t.TempDir()
	fn := filepath.Join(dir, "todo.json")

	if err := l.Save(fn); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	var l2 TodoList
	if err := l2.Get(fn); err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if len(l2) != 1 || l2[0].Task != "persist" {
		t.Fatalf("loaded data mismatch: %+v", l2)
	}

	// Get should not error on non-existent file
	var l3 TodoList
	if err := l3.Get(filepath.Join(dir, "no-such.json")); err != nil {
		t.Fatalf("expected no error for missing file, got %v", err)
	}

	// Empty file should be handled gracefully
	empty := filepath.Join(dir, "empty.json")
	if err := os.WriteFile(empty, []byte(""), 0644); err != nil {
		t.Fatalf("failed to create empty file: %v", err)
	}
	var l4 TodoList
	if err := l4.Get(empty); err != nil {
		t.Fatalf("expected no error for empty file, got %v", err)
	}
}
