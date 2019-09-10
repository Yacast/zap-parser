package zapparser

import (
	"path/filepath"
	"testing"
	"time"
)

func TestParser_OnClose(t *testing.T) {
	p, err := FromFile(filepath.Join("testdata", "10_logs"))
	if err != nil {
		t.Fatalf("failed loading test data : %s", err)
		return
	}

	called := false
	p.OnClose(func() {
		called = true
	})
	p.Start()
	if !called {
		t.Fatalf("OnClose has not been called")
	}

	p2, err := FromFile(filepath.Join("testdata", "10000_logs"))
	if err != nil {
		t.Fatalf("failed loading test data : %s", err)
		return
	}

	go func() {
		p2.Start()
	}()
	time.Sleep(10 * time.Millisecond)
	err = p2.OnClose(func() {})
	if err == nil {
		t.Fatalf("calling OnClose while parser running should have raised an error")
	}
}

func TestParser_OnError(t *testing.T) {
	p, err := FromFile(filepath.Join("testdata", "10000_logs"))
	if err != nil {
		t.Fatalf("failed loading test data : %s", err)
		return
	}

	expectedErrors := []string{
		"json parsing on line 12 failed: invalid character 'h' in literal true (expecting 'r')",
		"no timestamp in message at line 13",
		"unknown level at line 14: invalid level",
	}
	errCount := 0
	p.OnError(func(err error) {
		if errCount >= len(expectedErrors) {
			t.Fatalf("Unexpected error: %s", err)
		}
		if err.Error() != expectedErrors[errCount] {
			t.Fatalf(`Expected error "%s", got "%s"`, expectedErrors[errCount], err)
		}
		errCount++
	})
	p.Start()
	if errCount <= len(expectedErrors)-1 {
		t.Fatalf("Expected %d errors, go %d", len(expectedErrors), errCount)
	}

	p2, err := FromFile(filepath.Join("testdata", "10000_logs"))
	if err != nil {
		t.Fatalf("failed loading test data : %s", err)
		return
	}

	go func() {
		p2.Start()
	}()
	time.Sleep(10 * time.Millisecond)
	err = p2.OnError(func(err error) {})
	if err == nil {
		t.Fatalf("calling OnError while parser running should have raised an error")
	}
}

func TestParser_OnEntry(t *testing.T) {
	p, err := FromFile(filepath.Join("testdata", "10000_logs"))
	if err != nil {
		t.Fatalf("failed loading test data : %s", err)
		return
	}

	count := 0
	p.OnEntry(func(e *Entry) {
		count++
	})
	p.Start()
	if count != 10000 {
		t.Fatalf("Expected 10000 entries, go %d", count)
	}

	p2, err := FromFile(filepath.Join("testdata", "10000_logs"))
	if err != nil {
		t.Fatalf("failed loading test data : %s", err)
		return
	}

	go func() {
		p2.Start()
	}()
	time.Sleep(10 * time.Millisecond)
	err = p2.OnEntry(func(e *Entry) {})
	if err == nil {
		t.Fatalf("calling OnEntry while parser running should have raised an error")
	}
}
