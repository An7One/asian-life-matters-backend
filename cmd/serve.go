package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	api "github.com/zea7ot/web_api_aeyesafe/api"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "to start HTTP server with configured API",
	Long:  `to start a HTTP server and serve the configured API`,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := api.NewServer()
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// to define flags
	viper.SetDefault("port", "localhost:7777")
	viper.SetDefault("log_level", "debug")

}
