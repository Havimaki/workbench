package cmd

import (
	"fmt"
	"os"

	"github.com/havimaki/workbench/cmd/info"
	"github.com/havimaki/workbench/cmd/net"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "workbench",
	Short: "A golang workbench",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// ====== helpers
func setEnvDefaults() {
	viper.SetDefault("cmd.info.diskUsage.directory", ".")
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(info.InfoCmd)
}

func setConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".workbench" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".workbench")
	}

	// Checks if environment variables match any of the existing keys from
	// the config file, defaults or flags. If matching env vars are found,
	// they are loaded into Viper.
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// ====== handler
func init() {
	cobra.OnInitialize(setConfig)

	// Custom setup configs
	addSubcommandPalettes()
	setEnvDefaults()

	// Backup config file
	if err := viper.WriteConfigAs("workbench.default.yaml"); err != nil {
		fmt.Println(err)
	}

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.workbench.yaml)")
}
