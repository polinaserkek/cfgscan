package permissions

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckFilePermissions_Unsafe(t *testing.T) {
	dir := t.TempDir()

	path := filepath.Join(dir, "unsafe.yaml")

	err := os.WriteFile(path, []byte("test"), 0o600)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Chmod(path, 0o777)

	issues, err := CheckFilePermissions(path)

	if err != nil {
		t.Fatal(err)
	}

	if len(issues) != 2 {
		t.Fatalf("ожидалось 2 проблемы, получено %d", len(issues))
	}
}
func TestCheckFilePermissions_Safe(t *testing.T) {
	dir := t.TempDir()

	path := filepath.Join(dir, "safe.yaml")

	err := os.WriteFile(path, []byte("test"), 0o600)

	if err != nil {
		t.Fatal(err)
	}

	issues, err := CheckFilePermissions(path)

	if err != nil {
		t.Fatal(err)
	}

	if len(issues) != 0 {
		t.Fatalf("ожидалось 0 проблем, получено %d", len(issues))
	}
}
