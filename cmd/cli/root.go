package cli

import (
	"context"
	"fmt"
	"gojoke/internal/domain/app"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jedib0t/go-pretty/v6/text"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gojoke",
	Short: "CLI to get jokes from the D(e)ad jokes API",
	Long: `The GoJoke CLI is a command-line interface for fetching jokes from the D(e)ad jokes API.
With this CLI, you can easily get a random joke or search for jokes by keyword.

To use the CLI, simply run` + " `gojoke` " + `to get a random joke or` + " `gojoke find <query pattern>` " + `to search for jokes by keyword.
Refer to the help documentation for more information on the available commands and options.

Each subcommand supports a variety of options and flags that allow you to customize the behavior of the command.
For example, you can specify the number of jokes to fetch or the output format of the jokes.

The GoJoke CLI is designed to be easy to use and highly customizable.
Whether you're a developer looking for a quick laugh or a power user who needs to search for jokes on a regular basis, the GoJoke CLI has you covered.`,
	Version: "0.1.0",
	Run:     getRandomJoke,
}

func Execute(state app.State) {
	ctx := app.ContextWithState(state, rootCmd.Context())
	err := rootCmd.ExecuteContext(ctx)

	if err != nil && !strings.Contains(err.Error(), "unknown command") {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func getRandomJoke(cmd *cobra.Command, args []string) {
	ctx, err := app.StateFromContext(cmd.Context())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get app state: %v\n", err)
		os.Exit(1)
	}

	joke, err := ctx.JokeService().GetRandom()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get a random joke: %v\n", err)
		return
	}

	fmt.Printf("\n%s\n", text.Italic.Sprint(joke.Body))
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.SetContext(context.Background())
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $XDG_CONFIG_HOME/gojoke.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		configDir, err := os.UserConfigDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("gojoke")
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()
}
