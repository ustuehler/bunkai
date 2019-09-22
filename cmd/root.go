/*
Copyright © 2019 Uwe Stuehler <uwe@bsdx.de>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ustuehler/subs2srs/pkg/media"
	"github.com/ustuehler/subs2srs/pkg/subs"
)

var cfgFile string

var mediaFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "subs2srs <subsfile>",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		subsFile := args[0]

		subs, err := subs.OpenFile(subsFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: can't read subtitles: %v\n", err)
			os.Exit(1)
		}

		subsBase := strings.TrimSuffix(path.Base(subsFile), path.Ext(subsFile))
		outFile := path.Join(path.Dir(subsFile), subsBase+".tsv")
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

		for _, s := range subs.Items {
			if mediaFile != "" {
				audioFile, err := media.ExtractAudio(s.StartAt, s.EndAt, mediaFile, mediaPrefix)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: can't extract audio: %v\n", err)
					os.Exit(1)
				}

				imageFile, err := media.ExtractImage(s.StartAt, s.EndAt, mediaFile, mediaPrefix)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: can't extract image: %v\n", err)
					os.Exit(1)
				}

				startPos := media.TimePosition(s.StartAt)
				endPos := media.TimePosition(s.EndAt)
				if s.EndAt <= s.StartAt {
					endPos = ""
				}

				fmt.Fprintf(outStream, "%s\t%s\t%s\t<img src=\"%s\">\t[sound:%s]\t%s\n", subsBase, startPos, endPos, path.Base(imageFile), path.Base(audioFile), s)
			}
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

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.subs2srs.yaml)")
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
		viper.SetConfigName(".subs2srs")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
