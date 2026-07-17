package config

import (
	"os"
	"reflect"
	"testing"
)

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

func TestImportConfigFromFile_NotExists(t *testing.T) {
	tmp := t.TempDir()
	chdirForTest(t, tmp)

	cfg, err := ImportConfigFromFile()
	if err != nil {
		t.Fatalf("ImportConfigFromFile() error = %v", err)
	}

	if !reflect.DeepEqual(cfg, DefaultConfig()) {
		t.Fatalf("ImportConfigFromFile() = %#v, want %#v", cfg, DefaultConfig())
	}
}

func TestImportConfigFromFile_ValidJSON(t *testing.T) {
	tmp := t.TempDir()
	chdirForTest(t, tmp)

	data := []byte(`{"user_agent":"test-agent","data_path":"./cards.json","image_dir":"./img","image_base_path":"/assets","image_width":240,"rating":4,"page_title":"page","comment":"memo","release_date":"2026-07-17"}`)
	if err := os.WriteFile(configFile, data, 0o600); err != nil {
		t.Fatalf("WriteFile(%q) error = %v", configFile, err)
	}

	cfg, err := ImportConfigFromFile()
	if err != nil {
		t.Fatalf("ImportConfigFromFile() error = %v", err)
	}

	if cfg.UserAgent != "test-agent" {
		t.Errorf("UserAgent = %q, want %q", cfg.UserAgent, "test-agent")
	}
	if cfg.DataPath != "./cards.json" {
		t.Errorf("DataPath = %q, want %q", cfg.DataPath, "./cards.json")
	}
	if cfg.ImageDir != "./img" {
		t.Errorf("ImageDir = %q, want %q", cfg.ImageDir, "./img")
	}
	if cfg.ImageBasePath != "/assets" {
		t.Errorf("ImageBasePath = %q, want %q", cfg.ImageBasePath, "/assets")
	}
	if cfg.ImageWidth != 240 {
		t.Errorf("ImageWidth = %d, want %d", cfg.ImageWidth, 240)
	}
	if cfg.Rating != 4 {
		t.Errorf("Rating = %d, want %d", cfg.Rating, 4)
	}
	if cfg.PageTitle != "page" {
		t.Errorf("PageTitle = %q, want %q", cfg.PageTitle, "page")
	}
	if cfg.Comment != "memo" {
		t.Errorf("Comment = %q, want %q", cfg.Comment, "memo")
	}
	if cfg.ReleaseDate != "2026-07-17" {
		t.Errorf("ReleaseDate = %q, want %q", cfg.ReleaseDate, "2026-07-17")
	}
}

func TestImportConfigFromFile_InvalidJSON(t *testing.T) {
	tmp := t.TempDir()
	chdirForTest(t, tmp)

	if err := os.WriteFile(configFile, []byte("{"), 0o600); err != nil {
		t.Fatalf("WriteFile(%q) error = %v", configFile, err)
	}

	cfg, err := ImportConfigFromFile()
	if err == nil {
		t.Fatal("ImportConfigFromFile() error = nil, want non-nil")
	}

	if !reflect.DeepEqual(cfg, DefaultConfig()) {
		t.Fatalf("ImportConfigFromFile() = %#v, want %#v", cfg, DefaultConfig())
	}
}
