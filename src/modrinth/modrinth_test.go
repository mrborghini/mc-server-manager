package modrinth

import (
	"testing"
)

func TestInfoRequest(t *testing.T) {
	m := New()
	result, err := m.info.Get()
	if err != nil {
		t.Fatalf("info.get() returned an error: %v", err)
	}
	if result == nil {
		t.Fatal("info.get() returned nil result")
	}
	if result.Version == "" {
		t.Fatal("info.get().Version is empty")
	}
}
