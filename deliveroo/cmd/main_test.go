package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestParse_StepDoesNotLeakAcrossCommaSegments(t *testing.T) {
	got, err := parse("*/15,1-5", 0, 59)
	if err != nil {
		t.Fatalf("parse returned error: %v", err)
	}

	want := []int{0, 1, 2, 3, 4, 5, 15, 30, 45}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected result.\n got: %#v\nwant: %#v", got, want)
	}
}

func TestParse_InvalidStep(t *testing.T) {
	if _, err := parse("*/0", 0, 59); err == nil {
		t.Fatalf("expected error for invalid step, got nil")
	}
}

func TestParseCronExpression_CommandMayContainSpaces(t *testing.T) {
	_, _, _, _, _, cmd, err := parseCronExpression("*/15 0 1,15 * 1-5 /usr/bin/find -name foo")
	if err != nil {
		t.Fatalf("parseCronExpression returned error: %v", err)
	}
	if cmd != "/usr/bin/find -name foo" {
		t.Fatalf("unexpected command.\n got: %q\nwant: %q", cmd, "/usr/bin/find -name foo")
	}
}

func TestOutput_FieldNameIs14Columns(t *testing.T) {
	var buf bytes.Buffer
	output(&buf, []int{0, 15}, []int{0}, []int{1, 15}, []int{1, 2}, []int{1, 2, 3}, "/usr/bin/find")

	lines := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")
	if len(lines) != 6 {
		t.Fatalf("expected 6 lines, got %d", len(lines))
	}

	line := lines[0]
	if len(line) < 14 {
		t.Fatalf("expected at least 14 characters in first line, got %d: %q", len(line), line)
	}
	wantPrefix := "minute" + strings.Repeat(" ", 14-len("minute"))
	if line[:14] != wantPrefix {
		t.Fatalf("unexpected first 14 columns.\n got: %q\nwant: %q", line[:14], wantPrefix)
	}
}
