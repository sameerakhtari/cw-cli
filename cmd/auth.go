package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sameerakhtari/cw-cli/internal/cloudways"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with Cloudways API and save credentials",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter your Cloudways Email: ")
		emailInput, _ := reader.ReadString('\n')
		email := strings.TrimSpace(emailInput)

		fmt.Print("Enter your Cloudways API Key: ")
		apiKeyInput, _ := reader.ReadString('\n')
		apiKey := strings.TrimSpace(apiKeyInput)

		token, err := cloudways.GetAccessToken(email, apiKey)
		if err != nil {
			fmt.Println("❌ Login failed:", err)
			return
		}

		fmt.Println("🎉 Login successful! Access Token:", token)

		// Save config to file
		if err := cloudways.SaveConfig(email, apiKey, token); err != nil {
			fmt.Println("❌ Failed to save config:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
