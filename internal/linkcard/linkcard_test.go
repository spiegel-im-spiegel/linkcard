package linkcard

import (
	"context"
	"testing"

	"github.com/goark/webinfo"
	"github.com/spiegel-im-spiegel/linkcard/internal/config"
)

func TestNewLinkCard_RatingIsClampedWithoutMutatingConfig(t *testing.T) {
	cfg := &config.Config{ImageWidth: 240, Rating: 7, Comment: "memo"}
	info := &webinfo.Webinfo{
		URL:         "https://example.com",
		Title:       "Example Title",
		Description: "Example Description",
		ImageURL:    "https://example.com/image.png",
	}

	lc, err := newLinkCard(context.Background(), cfg, info)
	if err != nil {
		t.Fatalf("newLinkCard() error = %v", err)
	}

	if cfg.Rating != 7 {
		t.Fatalf("cfg.Rating = %d, want %d", cfg.Rating, 7)
	}
	if lc.ImageWidth != 240 {
		t.Fatalf("lc.ImageWidth = %d, want %d", lc.ImageWidth, 240)
	}
	if lc.Rating != 5 {
		t.Fatalf("lc.Rating = %d, want %d", lc.Rating, 5)
	}
	if lc.Comment != "memo" {
		t.Fatalf("lc.Comment = %q, want %q", lc.Comment, "memo")
	}
	if lc.Title != "Example Title" {
		t.Fatalf("lc.Title = %q, want %q", lc.Title, "Example Title")
	}

	for i := 0; i < 5; i++ {
		if !lc.Stars[i] {
			t.Fatalf("lc.Stars[%d] = false, want true", i)
		}
	}
}
