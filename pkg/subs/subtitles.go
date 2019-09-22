package subs

import (
	astisub "github.com/asticode/go-astisub"
)

// Subtitles represents a collection of subtitles corresponding to some media.
type Subtitles struct {
	*astisub.Subtitles
}

// OpenFile opens the given subtitles file for reading.
func OpenFile(filename string) (*Subtitles, error) {
	subs, err := astisub.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	return &Subtitles{subs}, nil
}
