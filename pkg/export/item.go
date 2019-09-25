package export

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	astisub "github.com/asticode/go-astisub"
	"github.com/ustuehler/bunkai/pkg/media"
	"github.com/ustuehler/bunkai/pkg/subs"
)

// ExportedItem represents the exported information of a single subtitle item,
// where Time is the primary field which identifies the item and ForeignCurr is
// the actual text of the item. The fields NativeCurr, NativePrev and NativeNext
// will be empty unless a second subtitle file was specified for the export and
// that second subtitle file is sufficiently aligned with the first.
type ExportedItem struct {
	Sound       string
	Time        string
	Source      string
	Image       string
	ForeignCurr string
	NativeCurr  string
	ForeignPrev string
	NativePrev  string
	ForeignNext string
	NativeNext  string
}

// ExportedItemWriter should write an exported item in whatever format is
// selected by the user.
type ExportedItemWriter func(*ExportedItem) error

// ExportItems calls the write function for each foreign subtitle item.
func ExportItems(foreignSubs, nativeSubs *subs.Subtitles, outputBase, mediaSourceFile, mediaPrefix string, write ExportedItemWriter) error {
	for i, foreignItem := range foreignSubs.Items {
		item, err := ExportItem(foreignItem, nativeSubs, outputBase, mediaSourceFile, mediaPrefix)
		if err != nil {
			return fmt.Errorf("can't export item #%d: %s: %v", i+1, foreignItem.String(), err)
		}

		if i > 0 {
			prevItem := foreignSubs.Items[i-1]
			item.ForeignPrev = prevItem.String()
		}

		if i+1 < len(foreignSubs.Items) {
			nextItem := foreignSubs.Items[i+1]
			item.ForeignNext = nextItem.String()
		}

		write(item)
	}
	return nil
}

func ExportItem(foreignItem *astisub.Item, nativeSubs *subs.Subtitles, subsBase, mediaFile, mediaPrefix string) (*ExportedItem, error) {
	item := &ExportedItem{}
	item.Source = subsBase
	item.ForeignCurr = joinLines(foreignItem.String())

	if nativeSubs != nil {
		if nativeItem := nativeSubs.Translate(foreignItem); nativeItem != nil {
			item.NativeCurr = joinLines(nativeItem.String())
		}
	}

	if mediaPrefix != "" {
		audioFile, err := media.ExtractAudio(foreignItem.StartAt, foreignItem.EndAt, mediaFile, mediaPrefix)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: can't extract audio: %v\n", err)
			os.Exit(1)
		}

		imageFile, err := media.ExtractImage(foreignItem.StartAt, foreignItem.EndAt, mediaFile, mediaPrefix)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: can't extract image: %v\n", err)
			os.Exit(1)
		}

		item.Time = timePosition(foreignItem.StartAt)
		item.Image = fmt.Sprintf("<img src=\"%s\">", path.Base(imageFile))
		item.Sound = fmt.Sprintf("[sound:%s]", path.Base(audioFile))
	}

	return item, nil
}

// timePosition formats the given time.Duration as a time code which can safely
// be used in file names on all platforms.
func timePosition(d time.Duration) string {
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	d -= s * time.Second
	ms := d / time.Millisecond
	return fmt.Sprintf("%02d:%02d:%02d,%03d", h, m, s, ms)
}

func joinLines(s string) string {
	s = strings.Replace(s, "\t", " ", -1)
	return strings.Replace(s, "\n", " ", -1)
}
