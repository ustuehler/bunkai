package media

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ffmpegExtractAudio(startAt, endAt time.Duration, inFile string, outFile string) error {
	inArgs := []string{
		"-ss", ffmpegPosition(startAt),
		"-i", inFile,
	}

	outArgs := []string{
		"-f", "mp3", "-vn",
		outFile,
	}
	if endAt > startAt {
		outArgs = append([]string{"-t", ffmpegPosition(endAt - startAt)}, outArgs...)
	}

	args := []string{
		"-loglevel", "error",
	}
	args = append(args, inArgs...)
	args = append(args, outArgs...)

	return ffmpeg(args...)
}

func ffmpegExtractImage(startAt, endAt time.Duration, inFile string, outFile string) error {
	var frameAt = startAt
	if endAt > startAt {
		frameAt = startAt + (endAt-startAt)/2
	}

	inArgs := []string{
		"-ss", ffmpegPosition(frameAt),
		"-i", inFile,
	}

	outArgs := []string{
		"-frames", "1",
		outFile,
	}
	if endAt > frameAt {
		outArgs = append([]string{"-t", ffmpegPosition(endAt - frameAt)}, outArgs...)
	}

	args := []string{
		"-loglevel", "error",
	}
	args = append(args, inArgs...)
	args = append(args, outArgs...)

	return ffmpeg(args...)
}

func ffmpegPosition(d time.Duration) string {
	s := d / time.Second
	d -= s * time.Second
	ms := d / time.Millisecond
	return fmt.Sprintf("%d.%d", s, ms)
}

func ffmpeg(arg ...string) error {
	cmd := exec.Command("ffmpeg", arg...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("ffmpeg command %v failed: %v", arg, err)
	}

	return nil
}
