package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "web_api_aeyesafe",
	Short: "A RESTful API for aeyesafe",
	Long:  "A RESTful API for aeyesafe",

	// Uncomment the following line if the bare application
	// has anything associated with it:
	// Run: func(cmd *cobra.Command, args[] string){}
}

// Execute adds all child commands to the root command and sets flags appropriately
// This is called by main.main().
// It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Cobra supports persistent flags,
	// which will be global for the entire application
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is &HOME/.aeyesafe_config.yaml")
	RootCmd.PersistentFlags().Bool("db_debug", false, "log sql to console")
	viper.BindPFlag("db_debug", RootCmd.PersistentFlags().Lookup("db_debug"))

	// Cobra also supports local flags,
	// which will only be executed when the action is called directly
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and environment variables if set
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// to find home directory
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		// to search config in home directory with name ".aeyesafe" (without extension)
		viper.AddConfigPath(home)
		viper.SetConfigName(".aeyesafe_config")
		viper.SetConfigType("yaml")
	}

	// to read in environment variables that match
	viper.AutomaticEnv()

	// to read it in if a config file has been found
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}
