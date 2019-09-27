package cmd

import "github.com/spf13/cobra"

var mediaFile string

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Decompose media for language study",
	Long: `The extract command group decomposes media into flash cards suitable
for studying a language, for example.`,
}

func init() {
	extractCmd.PersistentFlags().StringVarP(&mediaFile, "mediafile", "m", "", "media file to decompose")

	rootCmd.AddCommand(extractCmd)
}
