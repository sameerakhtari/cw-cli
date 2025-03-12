package cmd

import (
	"fmt"

	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/sameerakhtari/cw-cli/internal/cloudways"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Manage Cloudways servers",
}

var serverListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all servers",
	Run: func(cmd *cobra.Command, args []string) {
		servers, err := cloudways.ListServers()
		if err != nil {
			fmt.Println("❌ Failed to list servers:", err)
			return
		}

		if len(servers) == 0 {
			fmt.Println("No servers found.")
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Label", "IP", "Cloud", "Region", "Status"})

		for _, server := range servers {
			table.Append([]string{
				server.ID,
				server.Label,
				server.PublicIP,
				server.Cloud,
				server.Region,
				server.Status,
			})
		}

		table.Render()
	},
}

func init() {
	serverCmd.AddCommand(serverListCmd)
	rootCmd.AddCommand(serverCmd)
}
