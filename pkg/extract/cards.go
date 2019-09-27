package extract

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/ustuehler/bunkai/pkg/subs"
)

// ExtractCards specifies the parameters for generating a CSV export of the given
// foreign subtitles. The fields NativeSubtitles and MediaSource are optional.
type ExtractCards struct {
	ForeignSubtitlesFile string
	NativeSubtitlesFile  string
	MediaSourceFile      string
	OutputFieldSeparator string // defaults to "\t"
	OutputFileExtension  string // defaults to ".tsv" for "\t" and ".csv", otherwise
}

func (e *ExtractCards) setDefaults() {
	if e.OutputFieldSeparator == "" {
		e.OutputFieldSeparator = "\t"
	}

	if e.OutputFileExtension == "" {
		switch e.OutputFieldSeparator {
		case "\t":
			e.OutputFileExtension = ".tsv"
		default:
			e.OutputFileExtension = ".csv"
		}
	}
}

func (e *ExtractCards) outputBase() string {
	return strings.TrimSuffix(path.Base(e.ForeignSubtitlesFile), path.Ext(e.ForeignSubtitlesFile))
}

func (e *ExtractCards) outputFile() string {
	return path.Join(path.Dir(e.ForeignSubtitlesFile), e.outputBase()+"."+e.OutputFileExtension)
}

func (e *ExtractCards) mediaOutputDir() string {
	return path.Join(path.Dir(e.ForeignSubtitlesFile), e.outputBase()+".media")
}

func (e *ExtractCards) Execute() error {
	var nativeSubs *subs.Subtitles

	e.setDefaults()

	foreignSubs, err := subs.OpenFile(e.ForeignSubtitlesFile, false)
	if err != nil {
		return fmt.Errorf("can't read foreign subtitles: %v", err)
	}

	if e.NativeSubtitlesFile != "" {
		nativeSubs, err = subs.OpenFile(e.NativeSubtitlesFile, false)
		if err != nil {
			return fmt.Errorf("can't read native subtitles: %v", err)
		}
	}

	outStream, err := os.Create(e.outputFile())
	if err != nil {
		return fmt.Errorf("can't create output file: %s: %v", e.outputFile(), err)
	}
	defer outStream.Close()

	var mediaPrefix string
	if e.MediaSourceFile != "" {
		if err := os.MkdirAll(e.mediaOutputDir(), os.ModePerm); err != nil {
			return fmt.Errorf("can't create output directory: %s: %v", e.mediaOutputDir(), err)
		}
		mediaPrefix = path.Join(e.mediaOutputDir(), e.outputBase())
	}

	return ExportItems(foreignSubs, nativeSubs, e.outputBase(), e.MediaSourceFile, mediaPrefix, func(item *ExportedItem) error {
		fmt.Fprintf(outStream, "%s\t", item.Sound)
		fmt.Fprintf(outStream, "%s\t", item.Time)
		fmt.Fprintf(outStream, "%s\t", item.Source)
		fmt.Fprintf(outStream, "%s\t", item.Image)
		fmt.Fprintf(outStream, "%s\t", item.ForeignCurr)
		fmt.Fprintf(outStream, "%s\t", item.NativeCurr)
		fmt.Fprintf(outStream, "%s\t", item.ForeignPrev)
		fmt.Fprintf(outStream, "%s\t", item.NativePrev)
		fmt.Fprintf(outStream, "%s\t", item.ForeignNext)
		fmt.Fprintf(outStream, "%s\n", item.NativeNext)
		return nil
	})
}
