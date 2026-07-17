package facade

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/goark/gocli/rwi"
)

func decodeCardsForTest(t *testing.T, buf *bytes.Buffer) []map[string]any {
	t.Helper()

	var cards []map[string]any
	if err := json.Unmarshal(buf.Bytes(), &cards); err != nil {
		t.Fatalf("json.Unmarshal(output) error = %v, output = %s", err, buf.String())
	}
	if len(cards) != 1 {
		t.Fatalf("len(cards) = %d, want %d", len(cards), 1)
	}
	return cards
}

func chdirForTest(t *testing.T, dir string) {
	t.Helper()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd() error = %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("Chdir(%q) error = %v", dir, err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(wd); err != nil {
			t.Fatalf("restore working directory error = %v", err)
		}
	})
}

func TestRun_ReleaseDateFlagIsPropagatedToOutput(t *testing.T) {
	tmp := t.TempDir()
	chdirForTest(t, tmp)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<!doctype html><html><head><title>Example</title><meta name="description" content="Example Description"></head><body>ok</body></html>`))
	}))
	t.Cleanup(ts.Close)

	var out bytes.Buffer
	ui := rwi.New(rwi.WithWriter(&out), rwi.WithErrorWriter(&out))
	args := []string{"--release-date", "2026-07-17", ts.URL}

	if err := run(ui, args, NewVersionString("v0.0.0-test")); err != nil {
		t.Fatalf("run() error = %v", err)
	}

	cards := decodeCardsForTest(t, &out)

	got, ok := cards[0]["release_date"].(string)
	if !ok {
		t.Fatalf("release_date field missing or not string: %#v", cards[0]["release_date"])
	}
	if got != "2026-07-17" {
		t.Fatalf("release_date = %q, want %q", got, "2026-07-17")
	}
}

func TestRun_OtherFlagsArePropagatedToOutput(t *testing.T) {
	tmp := t.TempDir()
	chdirForTest(t, tmp)

	const expectedUA = "linkcard-test-agent"
	var gotUA string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUA = r.UserAgent()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<!doctype html><html><head><title>Original Title</title><meta name="description" content="Example Description"></head><body>ok</body></html>`))
	}))
	t.Cleanup(ts.Close)

	dataPath := filepath.Join(tmp, "cards.json")

	var out bytes.Buffer
	ui := rwi.New(rwi.WithWriter(&out), rwi.WithErrorWriter(&out))
	args := []string{
		"--user-agent", expectedUA,
		"--data-path", dataPath,
		"--image-width", "320",
		"--rating", "9",
		"--page-title", "CLI Title",
		"--comment", "memo-from-cli",
		ts.URL,
	}

	if err := run(ui, args, NewVersionString("v0.0.0-test")); err != nil {
		t.Fatalf("run() error = %v", err)
	}

	if gotUA != expectedUA {
		t.Fatalf("User-Agent = %q, want %q", gotUA, expectedUA)
	}

	cards := decodeCardsForTest(t, &out)
	card := cards[0]

	if got, _ := card["title"].(string); got != "CLI Title" {
		t.Fatalf("title = %q, want %q", got, "CLI Title")
	}
	if got, ok := card["image_width"].(float64); !ok || int(got) != 320 {
		t.Fatalf("image_width = %#v, want %d", card["image_width"], 320)
	}
	if got, ok := card["rating"].(float64); !ok || int(got) != 5 {
		t.Fatalf("rating = %#v, want %d (clamped)", card["rating"], 5)
	}
	if got, _ := card["comment"].(string); got != "memo-from-cli" {
		t.Fatalf("comment = %q, want %q", got, "memo-from-cli")
	}

	if _, err := os.Stat(dataPath); err != nil {
		t.Fatalf("saved data file stat error = %v", err)
	}
}
