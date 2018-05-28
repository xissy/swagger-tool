package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"os"

	"github.com/xissy/swagger-tool/swagger"
)

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge multiple Swagger JSON files to a JSON file",
	Long: `Merge multiple Swagger JSON files to a JSON file

Example: swagger-tool merge -i input1.json -i input2.json -o output.json

If -o option does not exist, it prints result on stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		inputFilenames, err := cmd.Flags().GetStringArray("input")
		exitIfErr(err)

		mergedSwaggerString, err := swagger.MergeSwaggerFiles(inputFilenames, nil)
		exitIfErr(err)

		outputFilename, err := cmd.Flags().GetString("output")
		exitIfErr(err)

		if outputFilename != "" {
			err := swagger.WriteFile(outputFilename, mergedSwaggerString)
			exitIfErr(err)
			return
		}

		fmt.Println(mergedSwaggerString)
	},
}

func exitIfErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(mergeCmd)

	mergeCmd.Flags().StringArrayP("input", "i", []string{}, "input Swagger JSON files (required)")
	mergeCmd.MarkFlagRequired("input")

	mergeCmd.Flags().StringP("output", "o", "", "output Swagger JSON file (prints stdout without -o)")
}
