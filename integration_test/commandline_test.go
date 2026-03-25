package integration_test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func moduleRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), ".."))
}

func testDataDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Join(filepath.Dir(file), "test_data")
}

func TestCLI_stdout_matches_expected_fixture(t *testing.T) {
	root := moduleRoot(t)
	fixtures := testDataDir(t)

	inputPath := filepath.Join(fixtures, "input.txt")
	expectedPath := filepath.Join(fixtures, "expected.txt")

	want, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("read expected fixture: %v", err)
	}

	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = root

	in, err := os.Open(inputPath)
	if err != nil {
		t.Fatalf("open input fixture: %v", err)
	}
	t.Cleanup(func() { _ = in.Close() })
	cmd.Stdin = in

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		t.Fatalf("go run main.go: %v\nstderr:\n%s", err, stderr.String())
	}
	if stderr.Len() > 0 {
		t.Fatalf("unexpected stderr:\n%s", stderr.String())
	}

	got := stdout.Bytes()
	if !bytes.Equal(got, want) {
		t.Fatalf("stdout mismatch: got %d bytes, want %d bytes", len(got), len(want))
	}
}
