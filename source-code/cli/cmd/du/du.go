package du

import (
	"fmt"
	"os"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/du_utils"
	"github.com/spf13/cobra"
)

var FlagNoHumanReadable bool
var FlagThreshold string
var FlagMaxDepth int

var Cmd = &cobra.Command{
	Use:   "du [<dir>]",
	Short: "Own implemetation of \"du\"",
	Args:  cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(args) == 1 {
			dir = args[0]
		}
		du_utils.RunDiskUsage(!FlagNoHumanReadable, FlagThreshold, dir, FlagMaxDepth, false)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagThreshold,
		"threshold",
		"t",
		"",
		"threshold of the size, any folders' size larger than the threshold will be print. for example, '1G', '10M', '100K', '1024'",
	)
	Cmd.Flags().BoolVarP(
		&FlagNoHumanReadable,
		"human-readable",
		"H",
		false,
		"disable human readable unit of size",
	)
	Cmd.Flags().IntVarP(
		&FlagMaxDepth,
		"max-depth",
		"d",
		0,
		"list its subdirectories and their sizes to any desired level of depth (i.e., to any level of subdirectories) in a directory tree.",
	)
}
