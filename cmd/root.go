package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jc0b/fleetdm-listener-go/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "fleetdm-listener",
	Short: "FleetDM Listener is an example of a multi-use binary for tracking FleetDM events, via PubSub or webhooks.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
		return initializeConfig(cmd)
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	PreRun:           util.PreRunSetup,
	TraverseChildren: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	pflags := rootCmd.PersistentFlags()

	pflags.Bool("debug", false, "Sets log level to debug. If multiple log level flags are set, the most verbose option will be respected.")
	viper.BindPFlag("debug", pflags.Lookup("debug"))

	pflags.Bool("trace", false, "Sets log level to trace. If multiple log level flags are set, the most verbose option will be respected.")
	viper.BindPFlag("trace", pflags.Lookup("trace"))

	pflags.Bool("json-logging", false, "Enables JSON logging when set.")
	viper.BindPFlag("json-logging", pflags.Lookup("json-logging"))
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	// When we bind flags to environment variables expect that the
	// environment variables are prefixed, e.g. a flag like --number
	// binds to an environment variable STING_NUMBER. This helps
	// avoid conflicts.
	v.SetEnvPrefix("LISTENER")

	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Bind to environment variables
	// Works great for simple config names, but needs help for names
	// like --favorite-color which we fix in the bindFlags function
	v.AutomaticEnv()

	// Bind the current command's flags to viper
	bindFlags(cmd, v)

	return nil
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name
		// If using camelCase in the config file, replace hyphens with a camelCased string.
		// Since viper does case-insensitive comparisons, we don't need to bother fixing the case, and only need to remove the hyphens.
		fmt.Printf("Name: %s\n", f.Name)
		fmt.Printf("Changed: %t\n", f.Changed)
		fmt.Printf("val: %s\n", f.Value.String())
		fmt.Printf("IsSet: %t\n", v.IsSet(configName))

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			fmt.Printf("%s: %s\n", configName, val)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
