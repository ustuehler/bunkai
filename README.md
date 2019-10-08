<a href="https://commons.wikimedia.org/wiki/File:Subtitleslogo.svg"><img src=".img/Subtitleslogo.svg" align="right" width="150px"/></a>

# Bunkai
[![Conventional Commits](.img/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)
> Decompose subtitles and media into personal language study material

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
- **Multiple subtitle formats**: Any format which is supported by [go-astisub][5]
  is also supported by this application, although some formats may work slightly
  better than others. If in doubt, try to use `.srt` subtitles.

[5]: https://github.com/asticode/go-astisub

## Installation
There is no proper release process at this time, nor a guarantee of stability
of any sort, as I'm the only user of the software that I am aware of. For now,
you must install the application from source.

Requirements:
- `go` command in `PATH` (only to build and install the application)
- `ffmpeg` command in `PATH` (used at runtime)

```bash
go get github.com/ustuehler/bunkai
```

## Usage
Bunkai is mainly used to generate flash cards from one or two subtitle files
and a corresponding media file.

For example:

```bash
bunkai extract cards -m media-content.mp4 foreign.srt native.srt
```

The above command generates the tab-separated file `foreign.tsv` and a
corresponding directory `foreign.media/` containing the associated images and
audio files. To do sentence mining, import the file `foreign.tsv` into a new
deck and then, at least in the case of Anki, copy the media files manually into
Anki's [collection.media directory](https://apps.ankiweb.net/docs/manual.html#file-locations).

Before you can import the deck with Anki though, you must add a new
[Note Type](https://apps.ankiweb.net/docs/manual.html#adding-a-note-type)
which includes some or all of the fields below on the front and/or back of
each card. The columns in the generated `.tsv` file are as follows:

| # | Name | Description |
| - | ---- | ----------- |
| 1 | Sound | Extracted audio as a `[sound]` tag for Anki |
| 2 | Time | Subtitle start time code as a string |
| 3 | Source | Base name of the subtitle source file |
| 4 | Image | Selected image frame as an `<img>` tag |
| 5 | ForeignCurr | Current text in foreign subtitles file |
| 6 | NativeCurr | Current text in native subtitles file |
| 7 | ForeignPrev | Previous text in foreign subtitles file |
| 8 | NativePrev | Previous text in native subtitles file |
| 9 | ForeignNext | Next text in foreign subtitles file |
| 10 | NativeNext | Next text in native subtitles file |

When you review the created deck for the first time, you should go quickly
through the entire deck at once. During this first pass, your goal should be
to identify those cards which you can understand almost perfectly, if not for
the odd piece of unknown vocabulary or grammar; all other cards which are
either too hard or too easy should be deleted in this pass. Any cards which
remain in the imported deck after mining should be refined and moved into your
regular deck for studying the language on a daily basis.

For other uses, run `bunkai --help` to view the built-in documentation.

## Subtitle editors
The state of affairs when it comes to open-source subtitle editors is a sad
one, but here's a list of editors which may or may not work passably. If you
know a good one, please let me know!

| Name | Platforms | Description |
| ---- | --------- | ----------- |
| [Aegisub](http://www.aegisub.org/) | macOS & others | Seems to have been a popular choice, but is no longer actively maintained. |
| [Jubler](https://www.jubler.org/) | macOS & others | Works reasonably well, but fixing timing issues is still somewhat cumbersome. |

## Known alternatives
There are at least two alternatives to this application that I know of.
Funny enough, I found substudy just after the initial prototype and might
not even have started Bunkai, had I seen it a bit earlier. :)

- [substudy](https://github.com/emk/subtitles-rs/tree/master/substudy): CLI alternative to subs2srs with the ability to export into other formats as well, not just SRS decks
- [subs2srs](http://subs2srs.sourceforge.net/): GUI software for Windows with many features, and inspiration for substudy and Bunkai

## Change log
See the file [CHANGELOG.md](CHANGELOG.md).

## License
[MIT](LICENSE)
