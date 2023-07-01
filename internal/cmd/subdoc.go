/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ziollek/cb-perf-tester/internal/cmd/utils"
	"github.com/ziollek/cb-perf-tester/internal/infra"
	"github.com/ziollek/cb-perf-tester/pkg/tester/model"
)

// subdocCmd represents the subdoc command
var subdocCmd = &cobra.Command{
	Use:   "subdoc",
	Short: "Test couchbase KV subdocument requests",
	Long:  `More info about subdocuments: https://docs.couchbase.com/java-sdk/current/howtos/subdocument-operations.html.`,
	Run: func(cmd *cobra.Command, args []string) {
		params := utils.ProvideSubDocParams(cmd)
		subDocBenchmark, err := infra.BuildSubDocBenchmark(params.Difficulty, params.SearchKeys)
		if err != nil {
			panic(fmt.Sprintf("Cannot prepare sub-document benchmark: %s", err))
		}
		doc := model.Generate("test-subdoc", params.Keys)
		fmt.Printf("benchmark params: %s\n", params.ToString())
		fmt.Printf("Generated doc with subkeys: %d, byte size is: %d\n\n", doc.Size, doc.JsonSize())
		fmt.Printf("subdoc report: %s\n", subDocBenchmark.DoExperiment(doc, params.Parallel, params.Repeats).ToString())
	},
}

func init() {
	subdocCmd.PersistentFlags().Int("keys", 10000, "define how many sub keys the tested document should consist of")
	subdocCmd.PersistentFlags().Int("search-keys", 4, "define how many sub keys should be seeked in single operation")
	subdocCmd.PersistentFlags().String("difficulty", "easy", "define how difficult is to find sub document [easy - first sub key, medium - in the middle, hard - on the end, impossible - try to find sub doc which does not exist]")

	rootCmd.AddCommand(subdocCmd)
}
