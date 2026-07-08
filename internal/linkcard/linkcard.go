package linkcard

import (
	"context"
	"crypto/sha1" // #nosec G505
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/goark/errs"
	"github.com/goark/webinfo"
	"github.com/spiegel-im-spiegel/linkcard/internal/config"
	"github.com/spiegel-im-spiegel/linkcard/internal/ecode"
)

// LinkCard represents the structure of a link card.
type LinkCard struct {
	HashID      string  `json:"hash_id"`
	URL         string  `json:"url"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	ImageURL    string  `json:"image_url,omitempty"`
	ImagePath   string  `json:"image_path,omitempty"`
	ImageWidth  int     `json:"image_width,omitempty"`
	Rating      int     `json:"rating,omitempty"`
	Stars       [5]bool `json:"stars,omitempty"`
	Comment     string  `json:"comment,omitempty"`
}

// newLinkCard creates a new LinkCard instance based on the provided configuration and web information.
func newLinkCard(ctx context.Context, cfg *config.Config, info *webinfo.Webinfo) (*LinkCard, error) {
	if cfg == nil || info == nil {
		return nil, ecode.ErrNullPointer
	}
	if info.URL == "" {
		return nil, errs.Wrap(ecode.ErrEmptyURL, errs.WithContext("page_url", info.URL))
	}

	// Create a new LinkCard instance based on the provided configuration, web information, and rating.
	lc := &LinkCard{
		Title:       info.Title,
		Description: info.Description,
		URL:         info.URL,
		ImageURL:    info.ImageURL,
		ImageWidth:  cfg.ImageWidth,
		Rating:      cfg.Rating,
		Comment:     cfg.Comment,
	}
	// Generate a unique hash ID (SHA1) for the link card based on the URL.
	h := sha1.Sum([]byte(info.URL)) // #nosec G401
	lc.HashID = hex.EncodeToString(h[:])
	// represent page title if cfg.PageTitle is not empty, otherwise use the title from webinfo
	if cfg.PageTitle != "" {
		lc.Title = cfg.PageTitle
	}

	// Fill in Stars based on the rating, ensuring it does not exceed 5.
	if lc.Rating > 0 {
		if lc.Rating > 5 {
			lc.Rating = 5
		}
		for i := range lc.Rating {
			lc.Stars[i] = true
		}
	}

	// Download the thumbnail image if ImageDir is specified in the configuration.
	if cfg.ImageDir != "" {
		tempf, err := info.DownloadThumbnail(ctx, cfg.ImageDir, cfg.ImageWidth, true)
		if err != nil {
			return nil, errs.Wrap(err, errs.WithContext("image_url", info.ImageURL))
		}
		if tempf == "" {
			return nil, errs.Wrap(ecode.ErrEmptyURL, errs.WithContext("image_url", info.ImageURL))
		}
		// Rename the downloaded image file to include the hash ID and maintain the original file extension.
		fname := filepath.Join(filepath.Dir(tempf), lc.HashID+filepath.Ext(tempf))
		if err := os.Rename(tempf, fname); err != nil {
			return nil, errs.Wrap(err, errs.WithContext("src_file", tempf), errs.WithContext("dst_file", fname))
		}
		base := filepath.Base(fname)
		lc.ImagePath = filepath.Join(cfg.ImageBasePath, base)
	}

	return lc, nil
}
