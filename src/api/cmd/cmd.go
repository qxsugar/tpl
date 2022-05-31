package cmd

import (
	"api-pkg/cmd/http"
	"api-pkg/internal/token"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "用法见 ./app --help",
	Run: func(cmd *cobra.Command, args []string) {
		http.Start()
	},
}

var api = &cobra.Command{
	Use:   "api",
	Short: "start api",
	Run: func(cmd *cobra.Command, args []string) {
		http.Start()
	},
}

var genToken = &cobra.Command{
	Use: "genToken",
	Run: func(cmd *cobra.Command, args []string) {
		tokenString, _ := token.NewToken(1)
		print(tokenString)
	},
}

func Execute() {
	rootCmd.AddCommand(api)
	rootCmd.AddCommand(genToken)
	_ = rootCmd.Execute()
}
