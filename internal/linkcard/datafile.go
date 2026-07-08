package linkcard

import (
	"encoding/json"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/goark/errs"
	"github.com/spiegel-im-spiegel/linkcard/internal/ecode"
)

// saveLinkCardToFile saves the provided link cards to a JSON file at the specified data path.
func saveLinkCardToFile(lc []LinkCard, dataPath string) error {
	if dataPath == "" || len(lc) == 0 { // If the data path is empty or there are no link cards, return without saving
		return nil
	}
	// Load existing link cards from the file
	cardsmap, err := loadLinkCardFromFile(dataPath, len(lc))
	if err != nil {
		return errs.Wrap(err, errs.WithContext("data_path", dataPath))
	}

	// Update the map with new link cards, using HashID as the key
	for _, card := range lc {
		cardsmap[card.HashID] = card
	}

	// Convert the map back to a slice for saving
	newlc := slices.SortedStableFunc(maps.Values(cardsmap), func(a, b LinkCard) int {
		return strings.Compare(a.HashID, b.HashID)
	})

	f, err := os.Create(filepath.Clean(dataPath))
	if err != nil {
		return errs.Wrap(err, errs.WithContext("data_path", dataPath))
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			err = errs.Join(err, errs.Wrap(cerr, errs.WithContext("data_path", dataPath)))
		}
	}()

	// Encode the link cards as JSON and write to the file
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(newlc); err != nil {
		return errs.Wrap(err, errs.WithContext("data_path", dataPath))
	}
	return nil
}

// loadLinkCardFromFile loads link cards from a JSON file at the specified data path and returns a map of LinkCard instances keyed by their HashID.
func loadLinkCardFromFile(dataPath string, bufsize int) (lcmap map[string]LinkCard, err error) {
	if dataPath == "" { // If the data path is empty, return an empty map and no error
		return nil, errs.Wrap(ecode.ErrInvalidDataPath, errs.WithContext("data_path", dataPath))
	}
	r, err := os.Open(filepath.Clean(dataPath))
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]LinkCard{}, nil
		}
		return nil, errs.Wrap(err, errs.WithContext("data_path", dataPath))
	}
	defer func() {
		if cerr := r.Close(); cerr != nil {
			err = errs.Join(err, errs.Wrap(cerr, errs.WithContext("data_path", dataPath)))
		}
	}()

	var cards []LinkCard
	if err := json.NewDecoder(r).Decode(&cards); err != nil {
		return nil, errs.Wrap(err, errs.WithContext("data_path", dataPath))
	}
	lcmap = make(map[string]LinkCard, len(cards)+bufsize)
	for _, card := range cards {
		lcmap[card.HashID] = card
	}
	return
}
