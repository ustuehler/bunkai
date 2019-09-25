package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ustuehler/bunkai/pkg/media"
	"github.com/ustuehler/bunkai/pkg/subs"
)

var cfgFile string

var mediaFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bunkai <subsfile1> [subsfile2]",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Args: argFuncs(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		var subsFile1, subsFile2 string
		var subs2 *subs.Subtitles
		var err error

		subsFile1 = args[0]

		if len(args) > 0 {
			subsFile2 = args[1]
			subs2, err = subs.OpenFile(subsFile2)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: can't read translated subtitles: %v\n", err)
				os.Exit(1)
			}
		}

		subs1, err := subs.OpenFile(subsFile1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: can't read original subtitles: %v\n", err)
			os.Exit(1)
		}

		subsBase := strings.TrimSuffix(path.Base(subsFile1), path.Ext(subsFile1))
		outFile := path.Join(path.Dir(subsFile1), subsBase+".tsv")
		outStream, err := os.Create(outFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "can't create output file: %s: %v", outFile, err)
			os.Exit(1)
		}

		mediaBase := strings.TrimSuffix(path.Base(mediaFile), path.Ext(mediaFile))
		mediaDir := path.Join(path.Dir(mediaFile), mediaBase+".media")
		mediaPrefix := path.Join(mediaDir, mediaBase)
		if err := os.MkdirAll(mediaDir, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "can't create output directory: %s: %v", mediaDir, err)
			os.Exit(1)
		}

		for _, item := range subs1.Items {
			mediaSource := subsBase
			expression := joinLines(item.String())
			fmt.Fprintf(outStream, "%s\t%s", mediaSource, expression)

			if subs2 != nil {
				expression2 := ""
				if item2 := subs2.Translate(item); item2 != nil {
					expression2 = joinLines(item2.String())
				}
				fmt.Fprintf(outStream, "\t%s", expression2)
			}

			if mediaFile != "" {
				leadTime := 100 * time.Millisecond
				startAt := item.StartAt - leadTime
				if startAt < 0 {
					startAt = item.StartAt
				}

				audioFile, err := media.ExtractAudio(startAt, item.EndAt, mediaFile, mediaPrefix)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: can't extract audio: %v\n", err)
					os.Exit(1)
				}

				imageFile, err := media.ExtractImage(item.StartAt, item.EndAt, mediaFile, mediaPrefix)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: can't extract image: %v\n", err)
					os.Exit(1)
				}

				meaning := fmt.Sprintf("<img src=\"%s\">", path.Base(imageFile))
				audio := fmt.Sprintf("[sound:%s]", path.Base(audioFile))
				fmt.Fprintf(outStream, "\t%s\t%s", meaning, audio)
			}

			fmt.Fprintf(outStream, "\n")
		}

		outStream.Close()
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
	rootCmd.PersistentFlags().StringVarP(&mediaFile, "mediafile", "m", "", "media file to extract audio from")
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

func joinLines(s string) string {
	s = strings.Replace(s, "\t", " ", -1)
	return strings.Replace(s, "\n", " ", -1)
}
