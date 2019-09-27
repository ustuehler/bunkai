package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ustuehler/bunkai/pkg/export"
)

var cfgFile string

var mediaFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bunkai <foreign-subs> [native-subs]",
	Short: "Decomposes media with subtitles into flash cards",
	Long: `Bunkai generates flash cards for an SRS like Anki from subtitles and
optional associated media content.

Example:
  bunkai -m media-content.mp4 foreign.srt native.srt

Based on the given subtitle files and associated media file, the above
command would create the tab-separated file "foreign.tsv" and a directory
"foreign.media/" containing images and audio files. Among other fields,
"foreign.tsv" would have a current, previous and next subtitle item from
both subtitle files, but the timing reference would be "foreign.srt".`,

	Args: argFuncs(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		var subsFile1, subsFile2 string

		subsFile1 = args[0]
		if len(args) > 1 {
			subsFile2 = args[1]
		}

		action := export.CSVExport{
			MediaSourceFile:      mediaFile,
			ForeignSubtitlesFile: subsFile1,
			NativeSubtitlesFile:  subsFile2,
			OutputFieldSeparator: "\t",
			OutputFileExtension:  "tsv",
		}

		if err := action.Execute(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.bunkai.yaml)")
	rootCmd.PersistentFlags().StringVarP(&mediaFile, "mediafile", "m", "", "media file associated with subtitles")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".subs2srs" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bunkai")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
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
