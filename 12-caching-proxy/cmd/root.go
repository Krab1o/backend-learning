package cmd

import (
	"fmt"
	"os"

	"caching-proxy/server"

	"github.com/spf13/cobra"
)

var port string
var url string

var rootCmd = &cobra.Command{
	Use:   "caching-proxy",
	Run: func(cmd *cobra.Command, args []string) {
	  server.Init(port, url)
	},
  }
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port on which application is running")
	rootCmd.Flags().StringVarP(&url, "origin", "o", "", "Site to make request via proxy")
	rootCmd.MarkFlagRequired("origin")
}