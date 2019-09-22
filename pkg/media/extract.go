package media

import (
	"fmt"
	"time"
)

// ExtractAudio copies the audio stream from the given media file in the given
// range to a file in the given output directory. The file name is generated
// automatically based on the given prefix, time range and audio format. If
// either startAt or endAt is 0, then the start or end of the media is assumed,
// accordingly.
func ExtractAudio(startAt, endAt time.Duration, inFile, outPrefix string) (string, error) {
	outFile := fmt.Sprintf("%s_%s-%s.mp3", outPrefix, pathPosition(startAt), pathPosition(endAt))
	return outFile, ffmpegExtractAudio(startAt, endAt, inFile, outFile)
}

// ExtractImage extracts a single frame from the given media file in the given
// time range. The image file name is generated automatically based on the
// given prefix, time range and audio format. If either startAt or endAt is 0,
// then the start or end of the media is assumed, accordingly.
func ExtractImage(startAt, endAt time.Duration, inFile, outPrefix string) (string, error) {
	outFile := fmt.Sprintf("%s_%s-%s.jpg", outPrefix, pathPosition(startAt), pathPosition(endAt))
	return outFile, ffmpegExtractImage(startAt, endAt, inFile, outFile)
}

// pathPosition formats the given time.Duration as a time code which can safely
// be used in file names on all platforms.
func pathPosition(d time.Duration) string {
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	d -= s * time.Second
	ms := d / time.Millisecond
	return fmt.Sprintf("%02dh%02dm%02ds%03dms", h, m, s, ms)
}
