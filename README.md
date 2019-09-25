# bunkai
> Decompose media subtitles into study material for language learning

Bunkai (分解), literally means "analysis" or "disassembly" in Japanese. The main purpose of the **bunkai** command-line tool is to generate flash cards for a Spaced Repetition System (SRS), e.g., Anki, from media content and corresponding subtitles. It is inspired by an article on [Sentence Mining](https://massimmersionapproach.com/table-of-contents/stage-1/jp-quickstart-guide/#sentence-mining), as well as existing tools (see [Known alternatives](#known-alternatives), below).

## Requirements
- `go` command in `PATH` (only at installation time)
- `ffmpeg` command in `PATH` (used at runtime)

## Installation
```bash
go get github.com/ustuehler/bunkai
```

## Change log
See the file [CHANGELOG.md](CHANGELOG.md).

## License
See the file [LICENSE](LICENSE).

## Known alternatives
- subs2srs
- substudy
