package facade

import "testing"

func TestNewVersionString(t *testing.T) {
	got := NewVersionString("v1.2.3").String()
	want := "linkcard v1.2.3\nhttps://github.com/spiegel-im-spiegel/linkcard"
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
