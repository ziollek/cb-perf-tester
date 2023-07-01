package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ziollek/cb-perf-tester/pkg/tester/model"
)

type GenericParams struct {
	Repeats  int
	Parallel int
}

func (gp *GenericParams) ToString() string {
	return fmt.Sprintf("repeats=%d, parallel=%d", gp.Repeats, gp.Parallel)
}

func ProvideParams(cmd *cobra.Command) *GenericParams {
	n, err := cmd.Flags().GetInt("repeat")
	if err != nil {
		panic(fmt.Sprintf("improper repeat option: %s", err))
	}
	goroutines, err := cmd.Flags().GetInt("parallel")
	if err != nil {
		panic(fmt.Sprintf("improper parallel option: %s", err))
	}

	return &GenericParams{
		Repeats:  n,
		Parallel: goroutines,
	}
}

type SubDocParams struct {
	Keys       int
	Difficulty model.Difficulty
	SearchKeys int
	*GenericParams
}

func ProvideSubDocParams(cmd *cobra.Command) *SubDocParams {
	keys, err := cmd.Flags().GetInt("keys")
	if err != nil {
		panic(fmt.Sprintf("improper keys option: %s", err))
	}
	searchKeys, err := cmd.Flags().GetInt("search-keys")
	if err != nil {
		panic(fmt.Sprintf("improper sub-keys option: %s", err))
	}
	difficultyName, err := cmd.Flags().GetString("difficulty")
	if err != nil {
		panic(fmt.Sprintf("improper difficulty option: %s", err))
	}

	return &SubDocParams{
		Keys:          keys,
		SearchKeys:    searchKeys,
		Difficulty:    model.FromString(difficultyName),
		GenericParams: ProvideParams(cmd),
	}
}

func (sp *SubDocParams) ToString() string {
	return fmt.Sprintf("keys=%d, level=%s, search-keys=%d, %s", sp.Keys, sp.Difficulty.ToString(), sp.SearchKeys, sp.GenericParams.ToString())
}

type RegularParams struct {
	Keys                 int
	SearchForNotExistent bool
	*GenericParams
}

func ProvideRegularParams(cmd *cobra.Command) *RegularParams {
	keys, err := cmd.Flags().GetInt("keys")
	if err != nil {
		panic(fmt.Sprintf("improper keys option: %s", err))
	}
	searchForNotExistent, err := cmd.Flags().GetBool("search-non-existent")
	if err != nil {
		panic(fmt.Sprintf("improper search-non-existent option: %s", err))
	}

	return &RegularParams{
		Keys:                 keys,
		SearchForNotExistent: searchForNotExistent,
		GenericParams:        ProvideParams(cmd),
	}
}

func (rp *RegularParams) ToString() string {
	return fmt.Sprintf("keys=%d, not-existent=%t, %s", rp.Keys, rp.SearchForNotExistent, rp.GenericParams.ToString())
}
