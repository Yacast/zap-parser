package zapparser

import (
	"path/filepath"
	"testing"
	"time"

	"go.uber.org/zap/zapcore"
)

func TestParser_Start(t *testing.T) {

	p, err := FromFile(filepath.Join("testdata", "10000_logs"))
	if err != nil {
		t.Fatalf("failed loading test data : %s", err)
		return
	}

	line := 1

	p.OnEntry(func(e *Entry) {
		if line == 10 {
			if e.Level != zapcore.WarnLevel {
				t.Fatalf("Line 10 must be Warn level")
			}
			if e.Message != "foo" {
				t.Fatalf("Line 10 message must be 'foo'")
			}
			if e.Timestamp.UnixNano() != time.Date(2019, 9, 10, 20, 11, 7, 40425062, time.UTC).UnixNano() {
				t.Fatalf("Line 10 timestamp must be '%s'", time.Date(2019, 9, 10, 20, 11, 7, 40425062, time.UTC).Format(time.RFC3339Nano))
			}
			if e.Caller != "zap-parser/parser_test.go:54" {
				t.Fatalf("Line 10 caller must be must be 'zap-parser/parser_test.go:54'")
			}
		}
		if line == 11 {
			if e.Level != zapcore.DPanicLevel {
				t.Fatalf("Line 11 must be DPanic level")
			}
			if e.Message != "bleeeh" {
				t.Fatalf("Line 11 message must be 'bleeeh'")
			}
			if e.Timestamp.UnixNano() != time.Date(2019, 9, 10, 20, 11, 7, 40429115, time.UTC).UnixNano() {
				t.Fatalf("Line 11 timestamp must be '%s'", time.Date(2019, 9, 10, 20, 11, 7, 40429115, time.UTC).Format(time.RFC3339Nano))
			}
			if e.Caller != "zap-parser/parser_test.go:48" {
				t.Fatalf("Line 11 caller must be must be 'zap-parser/parser_test.go:48'")
			}
			v, ok := e.Extras["v"]
			if !ok {
				t.Fatalf("Line 11 should have a 'v' extra field")
			}
			vv, ok := v.(float64)
			if !ok {
				t.Fatalf("Line 11 'v' extra field should be a float64")
			}
			if vv != 0.5391144752502441 {
				t.Fatalf("Line 11 'v' extra field should be `0.5391144752502441`")
			}
			v2, ok := e.Extras["attempt"]
			if !ok {
				t.Fatalf("Line 11 should have a 'attempt' extra field")
			}
			vv2, ok := v2.(float64)
			if !ok {
				t.Fatalf("Line 11 'attempt' extra field should be a float64")
			}
			if vv2 != 3 {
				t.Fatalf("Line 11 'attempt' extra field should be `3`")
			}
			v3, ok := e.Extras["stacktrace"]
			if !ok {
				t.Fatalf("Line 11 should have a 'stacktrace' extra field")
			}
			vv3, ok := v3.(string)
			if !ok {
				t.Fatalf("Line 11 'stacktrace' extra field should be a float64")
			}
			if vv3 != "github.com/Yacast/zap-parser.TestParser_Start\n\t/Users/jcrouzet/work/github.com/Yacast/zap-parser/parser_test.go:48\ntesting.tRunner\n\t/Users/jcrouzet/.gvm/gos/go1.12.5/src/testing/testing.go:865" {
				t.Fatalf("Line 11 'stacktrace' extra field should be `github.com/Yacast/zap-parser.TestParser_Start\n\t/Users/jcrouzet/work/github.com/Yacast/zap-parser/parser_test.go:48\ntesting.tRunner\n\t/Users/jcrouzet/.gvm/gos/go1.12.5/src/testing/testing.go:865`")
			}
		}
		if line == 12 {
			if e.Caller != "" {
				t.Fatalf("Line 12 should not have a caller field")
			}
		}

		line++
	})
	p.Start()
	if line != 10001 {
		t.Fatalf("10.000 logs should have been parsed, got %d", line-1)
	}
}

func BenchmarkParserFor10Entries_Start(b *testing.B) {
	for n := 0; n < b.N; n++ {
		p, err := FromFile(filepath.Join("testdata", "10_logs"))
		if err != nil {
			b.Fatalf("failed loading test data : %s", err)
			return
		}
		p.Start()
	}
}

func TestParser_Stop(t *testing.T) {

	p, err := FromFile(filepath.Join("testdata", "10000_logs"))
	if err != nil {
		t.Fatalf("failed loading test data : %s", err)
		return
	}

	line := 1

	p.OnEntry(func(e *Entry) {
		if line == 10 {
			p.Stop()
		}
		line++
	})
	p.Start()
	if line != 11 {
		t.Fatalf("10 logs should have been parsed, got %d", line-1)
	}

}
