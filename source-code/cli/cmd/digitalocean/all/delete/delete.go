package delete

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean/all"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete all resources in DigitalOcean account",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		log.Fatal("not implemented")
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
