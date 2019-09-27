package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ustuehler/bunkai/pkg/extract"
)

var extractCardsCmd = &cobra.Command{
	Use:   "cards <foreign-subs> [native-subs]",
	Short: "Decompose media into flash cards",
	Long: `This command generates flash cards for an SRS application like
Anki from subtitles and optional associated media content.

Example:
  bunkai extract cards -m media-content.mp4 foreign.srt native.srt

Based on the given subtitle files and associated media file, the above
command would create the tab-separated file "foreign.tsv" and a directory
"foreign.media/" containing images and audio files. Among other fields,
"foreign.tsv" would have a current, previous and next subtitle item from
both subtitle files, but the timing reference would be "foreign.srt".`,

	Args: argFuncs(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		var foreignSubs, nativeSubs string

		foreignSubs = args[0]

		if len(args) > 1 {
			nativeSubs = args[1]
		}

		action := extract.ExtractCards{
			ForeignSubtitlesFile: foreignSubs,
			NativeSubtitlesFile:  nativeSubs,
			MediaSourceFile:      mediaFile,
			OutputFieldSeparator: "\t",
			OutputFileExtension:  "tsv",
		}

		if err := action.Execute(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	extractCmd.AddCommand(extractCardsCmd)
}

// https://github.com/spf13/cobra/issues/648#issuecomment-393154805
func argFuncs(funcs ...cobra.PositionalArgs) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		for _, f := range funcs {
			err := f(cmd, args)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
