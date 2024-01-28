package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.Flags().StringP("manager", "m", "localhost:5555", "Manager to talk to")
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a running task.",
	Long: `cube stop command.

The stop command stops a running task.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		manager, _ := cmd.Flags().GetString("manager")
		url := fmt.Sprintf("http://%s/tasks/%s", manager, args[0])
		client := &http.Client{}
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			log.Printf("Error creating request %v: %v", url, err)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error connecting to %v: %v", url, err)
		}

		if resp.StatusCode != http.StatusNoContent {
			log.Printf("Error sending request: %v", err)
		}

		log.Printf("Task %v has been stopped.", args[0])
	},
}
