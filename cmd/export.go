package cmd

import "github.com/spf13/cobra"

var mediaFile string

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Decompose media for language study",
	Long: `The export command group decomposes media into flash cards suitable
for studying a language, for example.`,
}

func init() {
	exportCmd.PersistentFlags().StringVarP(&mediaFile, "mediafile", "m", "", "media file to decompose")

	rootCmd.AddCommand(exportCmd)
}
