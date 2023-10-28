package cli

import (
	"fmt"
	"gojoke/internal/domain/app"
	"gojoke/internal/ui"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Finds jokes that match provided pattern",
	Long: `The find command searches for jokes that match the specified search query.
You can search for jokes based on keywords or phrases, and the command will return a list of jokes that match the search query.
Note that the quality and appropriateness of the jokes returned by the API are outside of our control.
Use with caution and always double-check the jokes before sharing them.`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"query"},
	Run:       find,
}

func init() {
	rootCmd.AddCommand(findCmd)
}

func find(cmd *cobra.Command, args []string) {
	query := args[0]
	ctx, err := app.StateFromContext(cmd.Context())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get app state: %v\n", err)
		os.Exit(1)
	}

	jQuery, err := ctx.JokeService().Query(strings.TrimSpace(query))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to query jokes: %v\n", err)
		os.Exit(1)
	}

	jokes := make([]ui.ListItem, len(jQuery.Jokes))
	for i, j := range jQuery.Jokes {
		jokes[i] = j
	}

	item, err := ui.RunSelectList("Which one do you like the most?", jokes)
	if err != nil {
		switch e := err.(app.Error); e.Type {
		case app.Aborted:
			return
		}

		fmt.Fprintf(os.Stderr, "something went wrong :(\n")
		os.Exit(1)
	}

	fmt.Printf("\n%s\n", text.Italic.Sprint(item.TextFull()))
}
