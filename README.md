# Bunkai
> Decompose subtitles and media into personal language learning material

Bunkai (分解) literally means "analysis" or "disassembly" in Japanese. The Bunkai
application dissects subtitles and corresponding media files into flash cards
for [sentence mining][1] with an [SRS][2] system like [Anki][3]. It is inspired
by the linked article on sentence mining and [existing tools][4], which you
might want to check out as well.

[1]: https://massimmersionapproach.com/table-of-contents/stage-1/jp-quickstart-guide/#sentence-mining
[2]: https://en.wikipedia.org/wiki/Spaced_repetition
[3]: https://ankiweb.net/
[4]: #known-alternatives

## Features
- **One or two subtitle files**: Two subtitle files can be used together to
  provide foreign and native language expressions on the same card.
- **Media files are optional**: Requires only a single foreign subtitles file to
  generate text-only flash cards. Associated media content is optional, but
  highly recommended.

## Usage
Bunkai is mainly used to generate flash cards from one or two subtitle files
and a corresponding media file.

For example:

```bash
bunkai -m media-content.mp4 foreign.srt native.srt
```

The above command generates the tab-separated file `foreign.tsv` and a
corresponding directory `foreign.media/` containing the associated images and
audio files. To do sentence mining, import the file `foreign.tsv` into a new
deck and then, at least in the case of Anki, copy the media files manually into
Anki's media directory.

When you review the created deck, you should quickly go through all cards to
only those which you can almost understand, except for a single pice of unknown
vocabulary or grammar; all other cards which are either too hard or too easy
should be deleted in this pass. Any cards which remain in the imported deck
after mining should be refined and moved into your regular deck for studying the
language on a daily basis.

For other uses, run `bunkai --help` to view the built-in documentation.

## Installation
There are no official binary releases at this time. Until there are, you should
install the application from source.

Requirements:
- `go` command in `PATH` (only to build and install the application)
- `ffmpeg` command in `PATH` (used at runtime)

```bash
go get github.com/ustuehler/bunkai
```

## Change log
See the file [CHANGELOG.md](CHANGELOG.md).

## License
See the file [LICENSE](LICENSE).

## Known alternatives
- [substudy](https://github.com/emk/subtitles-rs/tree/master/substudy)
- [subs2srs](http://subs2srs.sourceforge.net/)
