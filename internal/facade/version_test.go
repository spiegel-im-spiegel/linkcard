package facade

import (
	"strings"
	"testing"
)

func TestNewVersionString(t *testing.T) {
	getGoVersion = func() string { return "(build with go1.27.3)" }
	got := NewVersionString("v1.2.3").String()
	want := "linkcard v1.2.3 (build with go1.27.3)\nhttps://github.com/spiegel-im-spiegel/linkcard"
	if got != want {
		t.Fatalf("NewVersionString().String() = %q, want %q", got, want)
	}
}

func TestUsageString(t *testing.T) {
	got := usageString()
	want := "linkcard\n\nUsage:\n  linkcard [flags] <url> [<url> ...]\n    <url> : URL(s) to generate link cards for\n\nFlags"
	if got != want {
		t.Fatalf("usageString() = %q, want %q", got, want)
	}
}

func TestReplaceVersion_WithExplicitVersion(t *testing.T) {
	getGoVersion = func() string { return "(build with go1.27.3)" }
	got := replaceVersion("v9.9.9")
	want := "v9.9.9 (build with go1.27.3)"
	if got != want {
		t.Fatalf("replaceVersion() = %q, want %q", got, want)
	}
}

func TestNewVersionString_EmptyVersionUsesFallback(t *testing.T) {
	got := NewVersionString("").String()
	lines := strings.Split(got, "\n")
	if len(lines) != 2 {
		t.Fatalf("version string line count = %d, want %d: %q", len(lines), 2, got)
	}
	if !strings.HasPrefix(lines[0], "linkcard ") {
		t.Fatalf("first line = %q, want prefix %q", lines[0], "linkcard ")
	}
	if lines[0] == "linkcard " {
		t.Fatal("first line has empty version payload")
	}
	if lines[1] != "https://github.com/spiegel-im-spiegel/linkcard" {
		t.Fatalf("second line = %q, want %q", lines[1], "https://github.com/spiegel-im-spiegel/linkcard")
	}
}

func TestJoinNonEmpty(t *testing.T) {
	got := joinNonEmpty("abc123", "", "(build with go1.26.5)")
	want := "abc123 (build with go1.26.5)"
	if got != want {
		t.Fatalf("joinNonEmpty() = %q, want %q", got, want)
	}
}
