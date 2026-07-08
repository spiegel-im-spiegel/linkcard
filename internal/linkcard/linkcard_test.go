package linkcard

import (
	"context"
	"testing"

	"github.com/goark/webinfo"
	"github.com/spiegel-im-spiegel/linkcard/internal/config"
)

func TestNewLinkCard_RatingIsClampedWithoutMutatingConfig(t *testing.T) {
	cfg := &config.Config{Rating: 7}
	info := &webinfo.Webinfo{
		URL: "https://example.com",
	}

	lc, err := newLinkCard(context.Background(), cfg, info)
	if err != nil {
		t.Fatalf("newLinkCard() error = %v", err)
	}

	if cfg.Rating != 7 {
		t.Fatalf("cfg.Rating = %d, want %d", cfg.Rating, 7)
	}

	for i := 0; i < 5; i++ {
		if !lc.Stars[i] {
			t.Fatalf("lc.Stars[%d] = false, want true", i)
		}
	}
}
