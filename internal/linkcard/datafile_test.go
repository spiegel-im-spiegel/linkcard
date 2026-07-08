package linkcard

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/spiegel-im-spiegel/linkcard/internal/ecode"
)

func TestLoadLinkCardFromFile_EmptyPath(t *testing.T) {
	_, err := loadLinkCardFromFile("", 1)
	if !errors.Is(err, ecode.ErrInvalidDataPath) {
		t.Fatalf("loadLinkCardFromFile() error = %v, want %v", err, ecode.ErrInvalidDataPath)
	}
}

func TestSaveLinkCardToFile_EmptyCardsNoop(t *testing.T) {
	tmp := t.TempDir()
	dataPath := filepath.Join(tmp, "cards.json")

	if err := saveLinkCardToFile([]LinkCard{}, dataPath); err != nil {
		t.Fatalf("saveLinkCardToFile() error = %v", err)
	}

	if _, err := os.Stat(dataPath); !os.IsNotExist(err) {
		t.Fatalf("file %q should not exist, stat error = %v", dataPath, err)
	}
}

func TestSaveLinkCardToFile_MergeAndSort(t *testing.T) {
	tmp := t.TempDir()
	dataPath := filepath.Join(tmp, "cards.json")

	seed := []LinkCard{
		{HashID: "b", URL: "https://old.example", Title: "old"},
	}
	// #nosec G304 -- dataPath is created from t.TempDir() in this test.
	f, err := os.Create(dataPath)
	if err != nil {
		t.Fatalf("Create(%q) error = %v", dataPath, err)
	}
	if err := json.NewEncoder(f).Encode(seed); err != nil {
		_ = f.Close()
		t.Fatalf("encode seed cards error = %v", err)
	}
	if err := f.Close(); err != nil {
		t.Fatalf("close seed file error = %v", err)
	}

	newCards := []LinkCard{
		{HashID: "a", URL: "https://a.example", Title: "new-a"},
		{HashID: "b", URL: "https://b.example", Title: "new-b"},
	}
	if err := saveLinkCardToFile(newCards, dataPath); err != nil {
		t.Fatalf("saveLinkCardToFile() error = %v", err)
	}

	// #nosec G304 -- dataPath is created from t.TempDir() in this test.
	r, err := os.Open(dataPath)
	if err != nil {
		t.Fatalf("Open(%q) error = %v", dataPath, err)
	}
	t.Cleanup(func() {
		_ = r.Close()
	})

	var saved []LinkCard
	if err := json.NewDecoder(r).Decode(&saved); err != nil {
		t.Fatalf("decode saved cards error = %v", err)
	}

	if len(saved) != 2 {
		t.Fatalf("len(saved) = %d, want %d", len(saved), 2)
	}
	if saved[0].HashID != "a" || saved[1].HashID != "b" {
		t.Fatalf("saved hash order = [%s, %s], want [a, b]", saved[0].HashID, saved[1].HashID)
	}
	if saved[1].URL != "https://b.example" || saved[1].Title != "new-b" {
		t.Fatalf("saved[1] = %#v, want updated card", saved[1])
	}
}

func TestLoadLinkCardFromFile_InvalidJSON(t *testing.T) {
	tmp := t.TempDir()
	dataPath := filepath.Join(tmp, "cards.json")
	if err := os.WriteFile(dataPath, []byte("{"), 0o600); err != nil {
		t.Fatalf("WriteFile(%q) error = %v", dataPath, err)
	}

	_, err := loadLinkCardFromFile(dataPath, 1)
	if err == nil {
		t.Fatal("loadLinkCardFromFile() error = nil, want non-nil")
	}
}
