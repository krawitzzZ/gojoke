package cli

import (
	"fmt"
	"gojoke/internal/infra/http"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Checks overall status of the app",
	Long: `Status command checks overall status of the app, its relevant resources and prints the info the stdout.
Currently it only checks the availability of the API.
More features will be added in the future, such as caching status, local database status, etc`,
	Run: func(cmd *cobra.Command, args []string) {
		if ok := http.GetStatus(); ok {
			fmt.Println("API is up and running!")
			return
		} else {
			fmt.Println("API is not reachable...")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
