package flow

import (
	"errors"
	"testing"
)

// An error in the step list reports its line in the file, counting the
// config header above the "---", not its line within the step document.
func TestParseErrorLineCountsConfigHeader(t *testing.T) {
	data := "appId: com.example\nname: x\n---\n- tapOn: \"A\"\n- notACommand: 1\n"
	_, err := Parse([]byte(data), "flow.yaml")
	var pe *ParseError
	if !errors.As(err, &pe) {
		t.Fatalf("err = %v, want *ParseError", err)
	}
	if pe.Line != 5 {
		t.Fatalf("line = %d, want 5 (%v)", pe.Line, err)
	}
}

func TestParseErrorLineWithoutHeader(t *testing.T) {
	_, err := Parse([]byte("- tapOn: \"A\"\n- notACommand: 1\n"), "flow.yaml")
	var pe *ParseError
	if !errors.As(err, &pe) || pe.Line != 2 {
		t.Fatalf("err = %v, want line 2", err)
	}
}
