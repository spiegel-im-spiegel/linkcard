package linkcard

import (
	"context"

	"github.com/goark/errs"
	"github.com/goark/webinfo"
	"github.com/spiegel-im-spiegel/linkcard/internal/config"
	"github.com/spiegel-im-spiegel/linkcard/internal/ecode"
)

// GenerateLinkCard generates link cards for the provided URLs based on the given configuration and saves them to a file.
func GenerateLinkCard(ctx context.Context, cfg *config.Config, urls []string) ([]LinkCard, error) {
	// Check if the configuration is nil, return an error if it is
	if cfg == nil {
		return nil, ecode.ErrNullPointer
	}
	size := len(urls)
	// Check if the list of URLs is empty, return an error if it is
	if size == 0 {
		return nil, ecode.ErrNoURL
	}
	cards := make([]LinkCard, 0, size)

	// Iterate over the list of URLs and generate link cards for each URL
	for _, u := range urls {
		lc, err := getWebInfo(ctx, cfg, u)
		if err != nil {
			return nil, errs.Wrap(err, errs.WithContext("url", u))
		}
		cards = append(cards, *lc)
	}
	if err := saveLinkCardToFile(cards, cfg.DataPath); err != nil {
		return nil, errs.Wrap(err, errs.WithContext("data_path", cfg.DataPath))
	}

	return cards, nil
}

// getWebInfo fetches web information for the given URL and creates a LinkCard instance based on the provided configuration.
func getWebInfo(ctx context.Context, cfg *config.Config, urlStr string) (*LinkCard, error) {
	info, err := webinfo.Fetch(ctx, urlStr, cfg.UserAgent)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", urlStr))
	}
	lc, err := newLinkCard(ctx, cfg, info)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", urlStr))
	}

	return lc, nil
}
