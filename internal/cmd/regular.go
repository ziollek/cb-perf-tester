/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/ziollek/cb-perf-tester/internal/cmd/utils"
	"github.com/ziollek/cb-perf-tester/internal/infra"
	"github.com/ziollek/cb-perf-tester/pkg/tester/model"

	"github.com/spf13/cobra"
)

// regularCmd represents the regular command
var regularCmd = &cobra.Command{
	Use:   "regular",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		params := utils.ProvideRegularParams(cmd)
		regularBenchmark, err := infra.BuildRegularBenchmark(params.SearchForNotExistent)
		if err != nil {
			panic(fmt.Sprintf("Cannot prepare regular benchmark: %s", err))
		}
		doc := model.Generate("test-regular", params.Keys)
		fmt.Printf("benchmark params: %s\n", params.ToString())
		fmt.Printf("Generated doc with subkeys: %d, byte size is: %d\n\n", doc.Size, doc.JsonSize())
		fmt.Printf("regular report: %s\n", regularBenchmark.DoExperiment(doc, params.Parallel, params.Repeats).ToString())
	},
}

func init() {
	regularCmd.PersistentFlags().Int("keys", 10000, "define how many sub keys the tested document should consist of")
	regularCmd.PersistentFlags().Bool("search-non-existent", false, "try to fetch non existent document")

	rootCmd.AddCommand(regularCmd)
}
